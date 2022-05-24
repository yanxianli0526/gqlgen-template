package migration

import (
	"graphql-go-template/internal/models"

	"gorm.io/gorm"
)

func UpdateMigration(db *gorm.DB) error {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	return db.AutoMigrate(
		&models.Category{},
		&models.Menu{},
	)
}
