package model

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Keywords string
	Content  string
}

type UserMenu struct {
	gorm.Model
	User     User `gorm:"ForeignKey:Uid"`
	Uid      uint `gorm:"not null"`
	Keywords string
	Content  string
}
