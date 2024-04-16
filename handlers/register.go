package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/steelthedev/danp/data"
	"github.com/steelthedev/danp/models"
)

func (h AppHandler) HanldeGetRegister(ctx *fiber.Ctx) error {

	context := fiber.Map{}
	return ctx.Render("sign-up", context)
}

func (h AppHandler) RegisterUser(ctx *fiber.Ctx) error {
	var userParams data.SignUpBody

	if err := ctx.BodyParser(&userParams); err != nil {
		return err
	}

	if errors := userParams.Validate(); len(errors) > 0 {
		fmt.Println(errors)
	}

	user := new(models.User)

	user.Email = userParams.Email
	user.Password = userParams.Password
	user.FirstName = userParams.FirstName
	user.LastName = userParams.LastName

	if result := h.DB.Create(&user); result.Error != nil {
		fmt.Println(result.Error)
	}
	return ctx.Redirect("/")
}
