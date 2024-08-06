package config

import (
	"gin-server/entity"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Enable the uuid-ossp extension
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	// Check if the enum type exists and create it if it doesn't
	createEnumType(db)

	db.AutoMigrate(&entity.Payment{})
	DB = db
}

func createEnumType(db *gorm.DB) {
	// Use raw SQL to check and create the enum type
	err := db.Exec(`
	DO $$
	BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status') THEN
					CREATE TYPE status AS ENUM ('CREATED', 'CONFIRMED', 'CANCELLED', 'DELIVERED');
			END IF;
	END $$;
	`).Error

	if err != nil {
		log.Fatal("failed to create enum type:", err)
	}
}
