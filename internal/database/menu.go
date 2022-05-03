package orm

import (
	"graphql-go-template/internal/models"

	"gorm.io/gorm"
)

func CreateMenu(db *gorm.DB, menu *models.Menu) error {
	return db.Create(menu).Error
}
