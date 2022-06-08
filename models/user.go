package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model

	// Username		string `gorm:"type:varchar(100);unique_index;not null;default:null" binding:"required"`
	Username		string `gorm:"type:varchar(100);unique_index;not null;default:null"`
}