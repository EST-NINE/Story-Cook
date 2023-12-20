package model

import (
	"gorm.io/gorm"
)

type Story struct {
	gorm.Model
	User     User   `gorm:"ForeignKey:Uid"`
	Uid      uint   `gorm:"not null"`
	Title    string `gorm:"not null"`
	Mood     string `gorm:"default:开心"`
	Keywords string
	Content  string `gorm:"type:longtext"`
}
