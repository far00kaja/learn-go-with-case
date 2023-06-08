package db

import (
	"log"
	"os"

	"github.com/far00kaja/learn-go-with-case/internal/auth/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	godotenv.Load()

	dbURL := os.Getenv("DB_CONNECT")

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	// dbURL := os.Getenv("DB_CONNECT")
	// dsn := fmt.Sprintf(
	// 	"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Jakarta",
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_NAME"),
	// )

	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(
		&models.Auth{},
	)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return db, nil

}
