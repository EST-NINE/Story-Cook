package model

import "gorm.io/gorm"

type UserCount struct {
	gorm.Model
	UID   uint `gorm:"not null"`
	Count int  `gorm:"default:1"`
	Date  string
}
