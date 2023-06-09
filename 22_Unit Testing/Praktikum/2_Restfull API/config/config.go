package config

import (
	"ORM_MVC/models"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitDBTest()
	InitialMigration()
	InitMigrateTest()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {

	config := Config{
		DB_Username: "root",
		DB_Password: "Nadiah271,.",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "praktikum1",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	InitialMigration()

}

func InitialMigration() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
}

type ConfigTest struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDBTest() {
	configTest := ConfigTest{
		DB_Username: "root",
		DB_Password: "Nadiah271,.",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "praktikum1_test",
	}

	connectionStringTest := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		configTest.DB_Username,
		configTest.DB_Password,
		configTest.DB_Host,
		configTest.DB_Port,
		configTest.DB_Name,
	)

	var err error
	DB, err = gorm.Open("mysql", connectionStringTest)
	if err != nil {
		panic(err)
	}
	InitMigrateTest()
}

func InitMigrateTest() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
}
