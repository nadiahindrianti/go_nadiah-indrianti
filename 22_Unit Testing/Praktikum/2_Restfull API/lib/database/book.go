package database

import (
	"ORM_MVC/config"
	"ORM_MVC/models"

	"github.com/labstack/echo"
)

func GetBooks() (interface{}, error) {
	var books []models.Book

	if e := config.DB.Find(&books).Error; e != nil {
		return nil, e
	}
	return books, nil
}

func GetBook(bookID int) (interface{}, error) {
	var book models.Book

	if err := config.DB.Where("id = ?", bookID).First(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func CreateBook(c echo.Context) (interface{}, error) {
	var book models.Book
	c.Bind(&book)

	if err := config.DB.Save(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func UpdateBook(c echo.Context) (interface{}, error) {
	var book models.Book

	if err := config.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		return nil, err
	}
	if err := c.Bind(&book); err != nil {
		return nil, err
	}
	config.DB.Save(&book)

	return book, nil
}

func DeleteBook(bookID int) (interface{}, error) {
	var book models.Book

	if err := config.DB.Where("id = ?", bookID).Delete(&book).Error; err != nil {
		return nil, err
	}

	return nil, nil
}
