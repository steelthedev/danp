package models

import "gorm.io/gorm"

const (
	Withdrawal string = "withdrawal"
	Deposit    string = "deposit"
)

type Transaction struct {
	gorm.Model

	ID     int64 `gorm:"primaryKey;autoIncrement"`
	UserId int64
	Amount float64 `gorm:"column:amount"`
	Type   string  `gorm:"column:type"`
	Status string  `gorm:"column:status"`
}
