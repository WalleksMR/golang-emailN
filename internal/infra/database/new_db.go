package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {
	dsn := "host=localhost user=docker password=docker dbname=docker port=54321 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("fail to connect to database")
	}

	return db
}
