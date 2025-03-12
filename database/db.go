package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=bachir password=rapido31 dbname=todo_go port=11002 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erreur de connexion à PostgreSQL :", err)
	}
	DB = db
	fmt.Println("📦 Connexion réussie à PostgreSQL !")
}
