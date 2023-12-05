package model

import "gorm.io/gorm"

type product struct {
	gorm.Model
	User        User   `gorm:"ForeignKey:Uid"`
	Uid         uint   `gorm:"not null"`
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
	Method      string `json:"method"`
}
