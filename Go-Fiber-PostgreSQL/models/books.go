package models

import "gorm.io/gorm"

type Book struct {
	ID        uint    `gorm:"primary_key;auto_increment" json:"id"`
	Title     *string `json:"title"`
	Author    *string `json:"author"`
	Publisher *string `json:"publisher"`
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Book{})
	return err
}
