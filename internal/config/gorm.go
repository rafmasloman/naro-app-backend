package config

import (
	"fmt"
	"log"
	"naro-app-be/internal/entity"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	DB_USER := os.Getenv(`DB_USER`)
	DB_NAME := os.Getenv(`DB_NAME`)
	DB_HOST := os.Getenv(`DB_HOST`)
	DB_PASSWORD := os.Getenv(`DB_PASSWORD`)
	DB_PORT := os.Getenv(`DB_PORT`)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)

	fmt.Println("dsn : ", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	MigrateDB(db)

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	return db

}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(
		&entity.User{},
	)
}
