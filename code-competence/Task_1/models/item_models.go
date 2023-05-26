package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Stock       string   `json:"stock"`
	Price       string   `json:"price"`
	CategoryID  uint     `json:"category_id"`
	Category    Category `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"category"`
}
