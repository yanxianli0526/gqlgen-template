package orm

import (
	"graphql-go-template/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateMenu(db *gorm.DB, menu *models.Menu) error {
	return db.Create(menu).Error
}

func GetMenuById(db *gorm.DB, menuId uuid.UUID) (*models.Menu, error) {
	var menu *models.Menu

	result := db.Where("id = ? ", menuId).First(&menu)

	return menu, result.Error
}

func GetMenus(db *gorm.DB) ([]*models.Menu, error) {
	var menus []*models.Menu

	result := db.Find(&menus)

	return menus, result.Error

}

func UpdateMenu(db *gorm.DB, menu *models.Menu) error {

	err := db.Model(&models.Menu{
		ID: menu.ID,
	}).Updates(map[string]interface{}{
		"item_name":       menu.ItemName,
		"price":           menu.Price,
		"is_stop_selling": menu.IsStopSelling,
	}).Error

	if err != nil {
		return err
	}
	return nil
}

func DeleteMenu(db *gorm.DB, menuId uuid.UUID) error {
	return db.Debug().Where("id = ?  ", menuId).Delete(&models.Menu{}).Error
}
