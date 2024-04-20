package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/steelthedev/danp/data"
	"github.com/steelthedev/danp/models"
	"github.com/sujit-baniya/flash"
)

func (h AppHandler) HandleGetDashboard(ctx *fiber.Ctx) error {
	user, err := h.GetAuthenticatedUser(ctx)
	if err != nil {
		return err
	}
	context := fiber.Map{
		"User": user,
	}
	return ctx.Render("dashboard/dashboard", context, "layout/base")
}

func (h AppHandler) HandleGetSettings(ctx *fiber.Ctx) error {
	user, err := h.GetAuthenticatedUser(ctx)
	if err != nil {
		return err
	}
	context := fiber.Map{
		"User": user,
	}

	return ctx.Render("dashboard/settings", context, "layout/base")
}

func (h AppHandler) EditUser(ctx *fiber.Ctx) error {
	var params data.EditUser

	if err := ctx.BodyParser(&params); err != nil {
		return flash.WithError(ctx, fiber.Map{"Error": "Invalid Body Request"}).Redirect("dashboard/settings")
	}

	err := params.CheckUserWithMailExists(ctx, h.DB)
	if err != nil {
		return flash.WithWarn(ctx, fiber.Map{"Error": err}).Redirect("/dashboard/settings")
	}

	user, err := h.GetAuthenticatedUser(ctx)
	if err != nil {
		return err
	}
	user.Email = params.Email
	user.FirstName = params.FirstName
	user.LastName = params.LastName
	user.Country = params.Country
	user.City = params.City
	user.PhoneNumber = params.PhoneNumber

	if result := h.DB.Save(&user); result.Error != nil {
		return flash.WithWarn(ctx, fiber.Map{"Error": "Could not update user on DB"}).RedirectBack("/dashboard/settings")
	}
	return flash.WithSuccess(ctx, fiber.Map{"Success": "Updates made successfully"}).Redirect("/dashboard/settings")

}

func (h AppHandler) HandleGetWallet(ctx *fiber.Ctx) error {
	user, err := h.GetAuthenticatedUser(ctx)
	if err != nil {
		return err
	}
	context := fiber.Map{
		"User": user,
	}
	return ctx.Render("dashboard/wallet", context, "layout/base")
}

func (h AppHandler) HandleGetInvestment(ctx *fiber.Ctx) error {
	var investments []models.Investment
	if result := h.DB.Find(&investments); result.Error != nil {
		return flash.WithWarn(ctx, fiber.Map{"Error": "An error occured"}).Redirect("/dashboard/investments")
	}
	context := fiber.Map{
		"Investments": investments,
	}
	return ctx.Render("dashboard/investments", context, "layout/base")
}
