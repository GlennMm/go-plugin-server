package utils

import (
	"gorm.io/gorm"
)

func DbInsert[T interface{}](db *gorm.DB, data *T) error {
	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&data).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func DbUpdate[T interface{}](db *gorm.DB, data *T) error {
	if err := db.Transaction(func(tx *gorm.DB) error {
		if err1 := tx.Save(&data).Error; err1 != nil {
			return err1
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
