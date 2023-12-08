package model

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Keywords string
	Content  string `gorm:"type:longtext"`
}
