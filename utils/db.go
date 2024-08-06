package utils

import (
	"test-eicon/config"
	"test-eicon/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDatabase(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.Order{})
	return db, nil
}
