package database

import (
	"example/jwt-authentication-tutorial-project/models"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(connectionString string) {
	Instance, dbError := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Failed to connect to database")
	}
	log.Println("Connected to database")
	
	Instance.AutoMigrate(&models.User{})
	log.Println("Migrated database")
}