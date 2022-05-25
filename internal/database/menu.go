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

func GetMenus(db *gorm.DB, allName []string) ([]*models.Menu, error) {
	var menus []*models.Menu

	// 可以用這個減低db的負擔 只針對需要的欄位去進行select(其實可做不可不做 真的做起來 可能會遇到其他麻煩)
	// 其實可以用下面的Find即可 並不會影響結果
	db.Debug().Select(allName).Find(&menus)
	// result := db.Find(&menus)

	return menus, nil

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
