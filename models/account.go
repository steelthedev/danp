package models

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string `gorm:"unique;not null"`
	FirstName   string `gorm:"column:first_name"`
	LastName    string `gorm:"column:last_name"`
	PhoneNumber int    `gorm:"column:phone_number"`
	Password    string `gorm:"not null"`
	Image       string `gorm:"column:image"`
	Wallet      Wallet
	Country     string `gorm:"column:country"`
	City        string `gorm:"column:city"`
}

func CheckUserExists(db *gorm.DB, email string) (*User, bool) {
	var user User
	if result := db.Where("Email=?", email).First(&user); result.Error == nil {
		return &user, true
	}
	return nil, false
}

func GetUserByID(db *gorm.DB, ID uint) (*User, error) {
	var user User
	if result := db.Where("ID=?", ID).First(&user); result.Error != nil {
		return nil, errors.New("could not get user")
	}
	return &user, nil
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	// create user wallet
	u.Wallet = Wallet{
		UserId:            u.ID,
		Balance:           0.00,
		TotalEarning:      0.00,
		PendingWithdrawal: 0.00,
	}
	return nil
}
