package database

import (
	"EmailGO/internal/campaign"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {

	dsn := "host=localhost user=postgres password=123456 dbname=emailGo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Fail to connect to db")
	}

	db.AutoMigrate(&campaign.Campaign{}, &campaign.Contact{})

	return db

}
