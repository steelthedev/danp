package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/steelthedev/danp/data"
	"github.com/steelthedev/danp/models"
	"github.com/steelthedev/danp/utils"
	"github.com/sujit-baniya/flash"
)

func (h AppHandler) HanldeGetRegister(ctx *fiber.Ctx) error {

	context := fiber.Map{}
	return ctx.Render("auth/sign-up", context)
}

func (h AppHandler) RegisterUser(ctx *fiber.Ctx) error {
	var userParams data.SignUpBody

	if err := ctx.BodyParser(&userParams); err != nil {
		return err
	}

	if errors := userParams.Validate(); len(errors) > 0 {
		return flash.WithError(ctx, errors).Redirect("/register")
	}

	user := new(models.User)

	user.Email = userParams.Email
	user.FirstName = userParams.FirstName
	user.LastName = userParams.LastName

	// Hash password
	hashedPwd, err := utils.HashPwd(userParams.Password)
	if err != nil {
		fmt.Println("An error occured")
	}
	user.Password = string(hashedPwd)

	if result := h.DB.Create(&user); result.Error != nil {
		data := fiber.Map{}
		data["error"] = "An error occured while creating user"
		return flash.RedirectToRoute(ctx, "/register", data, 500)
	}
	return ctx.Redirect("/login")
}

func (h AppHandler) HanldeGetLogin(ctx *fiber.Ctx) error {
	context := fiber.Map{}
	return ctx.Render("auth/sign-in", context)
}
