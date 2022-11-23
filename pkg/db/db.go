package db

import (
	"fmt"
	"log"

	"github.com/Antonio-Costa00/Go-Price-Monitor/pkg/credentials"
	"github.com/Antonio-Costa00/Go-Price-Monitor/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	c := credentials.GetCredentials()
	// dbURL := "postgres://admin:Passw0rd@localhost:5432/scraper"
	dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", c.DB_HOST, c.DB_USER, c.DB_PASS, c.DB, c.DB_PORT)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Product{})

	return db
}
