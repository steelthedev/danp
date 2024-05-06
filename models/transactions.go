package models

import "gorm.io/gorm"

const (
	Withdrawal string = "withdrawal"
	Deposit    string = "deposit"
)

type Transaction struct {
	gorm.Model
	UserId uint
	Amount float64 `gorm:"column:amount"`
	Type   string  `gorm:"column:type"`
	Status string  `gorm:"column:status"`
}
