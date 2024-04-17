package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/steelthedev/danp/data"
	"github.com/steelthedev/danp/models"
	"github.com/steelthedev/danp/utils"
	"github.com/sujit-baniya/flash"
	"golang.org/x/crypto/bcrypt"
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

func (h AppHandler) LoginUser(ctx *fiber.Ctx) error {
	var params data.Login
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}
	// check if user exists
	user, checkUser := models.CheckUserExists(h.DB, params.Email)
	if !checkUser {
		return flash.WithError(ctx, fiber.Map{
			"Error": "User does not exist",
		}).Redirect("/login")
	}
	// validate password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password)); err != nil {
		return flash.WithWarn(ctx, fiber.Map{"Error": "Please input correct password and try again"}).Redirect("/login")
	}

	// start log in user
	sesssion, err := h.Store.Get(ctx)
	if err != nil {
		fmt.Println("Could not get session")
	}
	sesssion.Set("userID", user.ID)
	sesssion.Set("isAuthenticated", true)
	sesssion.SetExpiry(time.Second * 40000000)
	sesssion.Save()

	return ctx.Redirect("/dashboard")

}

func (h AppHandler) GetAuthenticatedUser(ctx *fiber.Ctx) (*models.User, error) {
	session, err := h.Store.Get(ctx)
	if err != nil {
		return nil, err
	}
	userID := session.Get("userID")
	if userID == nil {
		return nil, err
	}
	user, err := models.GetUserByID(h.DB, userID.(int64))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (h AppHandler) Logout(ctx *fiber.Ctx) error {

	session, err := h.Store.Get(ctx)
	if err != nil {
		return err
	}
	if err := session.Destroy(); err != nil {
		return err
	}

	return ctx.Redirect("/login")
}
