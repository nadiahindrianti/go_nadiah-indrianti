package database

import (
	"Task_1/configs"
	"Task_1/models"
)

func CreateItem(i models.Item) (interface{}, error) {
	if err := configs.DB.Create(&i).Error; err != nil {
		return nil, err
	}

	if err := configs.DB.Joins("Category").Find(&i).Error; err != nil {
		return nil, err
	}
	return i, nil
}

func UpdateItemController(itemID uint, i models.Item) (interface{}, error) {
	item := models.Item{}
	item.ID = itemID
	if err := configs.DB.Joins("Category").Find(&item).Error; err != nil {
		return nil, err
	}

	item.Name = i.Name
	item.Description = i.Description
	item.Price = i.Price
	item.Stock = i.Stock
	item.CategoryID = i.CategoryID

	err := configs.DB.Save(&item).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

func GetItembyCategory(Category int) (interface{}, error) {
	var item []models.Item

	if err := configs.DB.Joins("Category").Find(&item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func GetItembyName(itemName string) (interface{}, error) {
	var item models.Item
	item.Name = itemName

	if err := configs.DB.First(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func GetItemControllerbyId(itemID uint) (interface{}, error) {
	var item models.Item
	item.ID = itemID

	if err := configs.DB.First(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func GetItemController() (interface{}, error) {
	var itemAll []models.Item
	if err := configs.DB.Find(&itemAll).Error; err != nil {
		return nil, err
	}

	return itemAll, nil
}

func DeleteItemController(itemID int) (interface{}, error) {
	err := configs.DB.Delete(&models.Item{}, itemID).Error
	if err != nil {
		return nil, err
	}
	return itemID, nil
}
