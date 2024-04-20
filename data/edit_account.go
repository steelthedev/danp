package data

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/steelthedev/danp/models"
	"gorm.io/gorm"
)

type EditUser struct {
	Email       string
	FirstName   string
	LastName    string
	PhoneNumber int
	Image       []byte
	Country     string
	City        string
}

func (h EditUser) CheckUserWithMailExists(ctx *fiber.Ctx, db *gorm.DB) error {

	_, ok := models.CheckUserExists(db, h.Email)

	if ok {
		return fmt.Errorf("user with this email exists")
	}
	return nil
}
