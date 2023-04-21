package database

import (
	"ORM_MVC/config"
	"ORM_MVC/models"
)

func GetBooks() (interface{}, error) {
	var books []models.Book

	if e := config.DB.Find(&books).Error; e != nil {
		return nil, e
	}
	return books, nil
}
