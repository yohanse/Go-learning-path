package database

import (
	"example/jwt-authentication-tutorial-project/models"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func Connect(connectionString string) {
	var dbError error
	Instance, dbError = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Failed to connect to database")
	}
	log.Println("Connected to database")
	
	Instance.AutoMigrate(&models.User{})
	log.Println("Migrated database")
}