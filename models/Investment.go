package models

import "gorm.io/gorm"

type Investment struct {
	gorm.Model
	Title         string  `gorm:"column:title"`
	Price         float64 `gorm:"column:price"`
	Roi           float64 `gorm:"column:roi"`
	ReferralBonus float64 `gorm:"column:bonus"`
}
