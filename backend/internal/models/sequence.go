package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Sequence struct {
	gorm.Model
	ID       uuid.UUID `json:"id" gorm:"type:uuid"`
	GotchiID uuid.UUID `json:"gotchi_id" gorm:"type:uuid"`
	Expires  time.Time `json:"expires" gorm:"not null"`
	Sequence string    `json:"sequence" gorm:"not null"`
}

func (s *Sequence) BeforeCreate(db *gorm.DB) (err error) {
	s.ID = uuid.New()

	// Delete all existing sequences
	err = db.Unscoped().Where("gotchi_id = ?", s.GotchiID).Delete(&Sequence{}).Error

	if err != nil {
		return err
	}

	s.Expires = time.Now().Add(60 * time.Second)

	return nil
}

func (s *Sequence) Create(db *gorm.DB) (Sequence, error) {
	result := db.Create(&s)

	if result.Error != nil {
		return *s, result.Error
	}

	return *s, nil
}
