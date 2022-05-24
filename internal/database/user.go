package orm

import (
	"graphql-go-template/internal/models"

	"gorm.io/gorm"
)

func GetUsers(db *gorm.DB) ([]*models.User, error) {
	var users []*models.User

	result := db.Find(&users)

	return users, result.Error

}
