package models

import (
	"math"

	"gorm.io/gorm"
)

type Investment struct {
	gorm.Model
	UserId             uint
	Title              string  `gorm:"column:title"`
	Price              float64 `gorm:"column:price"`
	PercentageIncrease float64 `gorm:"column:percentage_increase"`
	Roi                float64 `gorm:"column:roi"`
	ReferralBonus      float64 `gorm:"column:bonus"`
	FinalAmount        float64 `gorm:"column:final_amount"`
	Duration           int16   `gorm:"column:duration"`
}

func (h Investment) CalcFinalAmount() float64 {
	return h.Price * math.Pow(1+h.PercentageIncrease, float64(h.Duration))
}

func (h Investment) CalcRoi() float64 {
	finalAmount := h.CalcFinalAmount()
	roi := ((finalAmount - h.Price) / h.Price) * 100
	return roi
}

func (h *Investment) BeforeCreate(db *gorm.DB) error {
	h.Roi = h.CalcRoi()
	h.FinalAmount = h.CalcFinalAmount()
	return nil
}
