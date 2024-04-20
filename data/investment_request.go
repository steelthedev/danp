package data

type Investment struct {
	Title              string  `gorm:"form:title"`
	Price              float64 `gorm:"form:price"`
	PercentageIncrease float64 `gorm:"form:percentage_increase"`
	ReferralBonus      float64 `gorm:"form:bonus"`
	Duration           int16   `gorm:"form:duration"`
}
