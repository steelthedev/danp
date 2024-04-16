package models

type User struct {
	ID          int64  `gorm:"primaryKey;autoIncrement"`
	Email       string `gorm:"unique;not null"`
	FirstName   string `gorm:"column:first_name"`
	LastName    string `gorm:"column:last_name"`
	PhoneNumber int    `gorm:"column:phone_number"`
	Password    string `gorm:"not null"`
	Image       []byte `gorm:"type:bytea"`
}
