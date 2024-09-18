package database

import (
	"EmailGO/internal/campaign"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {

	dsn := os.Getenv("DATABASE")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Fail to connect to db")
	}

	db.AutoMigrate(&campaign.Campaign{}, &campaign.Contact{})

	return db

}
