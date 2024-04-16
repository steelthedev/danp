package data

import (
	"github.com/gofiber/fiber/v2"
	"github.com/steelthedev/danp/utils"
)

type SignUpBody struct {
	Email       string
	FirstName   string
	LastName    string
	PhoneNumber int
	Password    string
}

func (p SignUpBody) Validate() fiber.Map {
	data := fiber.Map{}

	if len(p.FirstName) < 1 {
		data["firstNameError"] = "First name can not be empty"
	}

	if len(p.LastName) < 1 {
		data["lastNameError"] = "last name can not be empty"
	}

	if !utils.IsValidEmail(p.Email) {
		data["emailError"] = "Invalid Email Address"
	}

	if len(p.Password) < 6 {
		data["passwordError"] = "Password must be longer than 6 characters"
	}

	return data

}
