package models

import "gorm.io/gorm"

type Files struct {
	CommonModelFields
	Title       string `json:"title"`
	Description string `json:"description"`
	File        string `json:"musicFile"`
	OwnerID     uint   `json:"ownerId"`
	Owner       User   `json:"owner"`
}

func ManuscriptMigrate(db *gorm.DB) {
	db.AutoMigrate(&Files{})
}
