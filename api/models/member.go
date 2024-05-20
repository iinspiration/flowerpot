package models

import (
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Name         string
	Lastname     string
	MobileNumber string
	Email        string
	Address      string
	TaxDetail    string
	Level        string `gorm:"type:ENUM('normal', 'regular', 'vip')"`
	Status       string `gorm:"type:ENUM('active', 'inactive')"`
}
