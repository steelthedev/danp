package models

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	UserId uint
	// User              User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Balance           float64
	TotalEarning      float64
	PendingWithdrawal float64
}
