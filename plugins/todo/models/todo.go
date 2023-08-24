package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Name        string `gorm:"name;index"`
	Description string `gorm:"description"`
	Done        bool   `gorm:"done"`
}
