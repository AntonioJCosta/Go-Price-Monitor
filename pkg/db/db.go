package db

import (
	"fmt"
	"log"

	"github.com/Antonio-Costa00/Go-Price-Monitor/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(host, user, password, dbName, port string) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.Product{})
	return db
}
