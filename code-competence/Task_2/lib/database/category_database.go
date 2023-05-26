package database

import (
	"Task_2/configs"
	"Task_2/models"
)

func CreateCategoryController(category models.Category) (interface{}, error) {
	if err := configs.DB.Create(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func UpdateCategoryController(CategoryID uint, categoryupdate models.Category) (interface{}, error) {
	category := models.Category{}
	category.ID = CategoryID

	configs.DB.First(&category)

	category.Name = categoryupdate.Name

	err := configs.DB.Save(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func DeleteCategoryController(CategoryID int) (interface{}, error) {
	if err := configs.DB.Delete(&models.Category{}, CategoryID).Error; err != nil {
		return nil, err
	}
	return CategoryID, nil
}

func GetCategoryControllerAll() (interface{}, error) {
	var category []models.Category

	if err := configs.DB.Find(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func GetCategoryController(id int) (interface{}, error) {
	var category models.Category

	if err := configs.DB.First(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}
