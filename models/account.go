package models

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          int64  `gorm:"primaryKey;autoIncrement"`
	Email       string `gorm:"unique;not null"`
	FirstName   string `gorm:"column:first_name"`
	LastName    string `gorm:"column:last_name"`
	PhoneNumber int    `gorm:"column:phone_number"`
	Password    string `gorm:"not null"`
	Image       []byte `gorm:"type:bytea"`
}

func CheckUserExists(db *gorm.DB, email string) (*User, bool) {
	var user User
	if result := db.Where("Email=?", email).First(&user); result.Error == nil {
		return &user, true
	}
	return nil, false
}

func GetUserByID(db *gorm.DB, ID int64) (*User, error) {
	var user User
	if result := db.Where("ID=?", ID).First(&user); result.Error != nil {
		return nil, errors.New("could not get user")
	}
	return &user, nil
}
