package db

import (
	"redler/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	m := []interface{}{
		&models.UserModel{},
	}
	return db.AutoMigrate(m...)
}
