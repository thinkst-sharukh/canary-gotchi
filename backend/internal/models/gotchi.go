package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Gotchi struct {
	gorm.Model
	ID        uuid.UUID `json:"id" gorm:"type:uuid"`
	Name      string    `json:"name" gorm:"not null" validate:"required,min=3,max=32"`
	Hash      string    `json:"hash" gorm:"not null" validate:"required"`
	AuthToken string    `json:"auth_token" gorm:"not null" validate:"required"`
	Level     int       `json:"level" gorm:"not null"`
	Verified  bool      `json:"verified" gorm:"not null;default:false"`
	Sequence  Sequence  `json:"sequence" gorm:"foreignKey:GotchiID;constraint:OnDelete:CASCADE"`
}

func (g *Gotchi) Create(db *gorm.DB) (Gotchi, error) {

	// check if gotchi already exists
	err := db.Preload("Sequence", "expires > ?", time.Now()).First(&g, "id = ?", g.ID).Error

	if err == nil {
		return *g, fmt.Errorf("Duplicate enrollment can't be created")
	}

	result := db.Create(&g)

	if result.Error != nil {
		return *g, result.Error
	}

	return *g, nil
}

func (g *Gotchi) Get(db *gorm.DB) (Gotchi, error) {
	err := db.First(&g, "id = ?", g.ID).Error

	return *g, err
}

func (g *Gotchi) GetWithSequence(db *gorm.DB) (Gotchi, error) {
	err := db.Preload("Sequence", "expires > ?", time.Now()).First(&g, "id = ?", g.ID).Error

	return *g, err
}

func (g *Gotchi) Save(db *gorm.DB) error {
	err := db.Save(&g).Error

	return err
}
