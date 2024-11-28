package database

import (
	"backend/internal/models"
	"backend/internal/utils"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	database = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
	schema   = os.Getenv("DB_SCHEMA")
)

func New() *gorm.DB {
	isDev := utils.IsDev()
	sslMode := "disable"

	// if !isDev {
	// 	// sslMode = "verify-full"
	// 	sslMode = "require"
	// }

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&search_path=%s", username, password, host, port, database, sslMode, schema)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if err != nil {
		panic(err)
	}

	log.Println("Connected")

	if isDev {
		db.Logger = logger.Default.LogMode(logger.Info)
	} else {
		db.Logger = logger.Default.LogMode(logger.Silent)
	}

	// db.Migrator().DropTable(&models.Gotchi{}, &models.Sequence{})

	db.AutoMigrate(&models.Gotchi{}, &models.Sequence{})

	return db
}
