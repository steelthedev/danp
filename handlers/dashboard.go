package handlers

import (
	"log/slog"

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
		return flash.WithError(ctx, fiber.Map{"message": "Invalid Body Request"}).Redirect("/dashboard/settings")
	}

	user, err := h.GetAuthenticatedUser(ctx)
	if err != nil {
		return err
	}
	user.FirstName = params.FirstName
	user.LastName = params.LastName
	user.Country = params.Country
	user.City = params.City
	user.PhoneNumber = params.PhoneNumber

	if result := h.DB.Save(&user); result.Error != nil {
		slog.Info("Error updating user details: %v", result.Error)
		return flash.WithWarn(ctx, fiber.Map{"message": "Could not update user on DB"}).RedirectBack("/dashboard/settings")
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
		return flash.WithWarn(ctx, fiber.Map{"message": "An error occured"}).Redirect("/dashboard/investments")
	}
	context := fiber.Map{
		"Investments": investments,
	}
	return ctx.Render("dashboard/investments", context, "layout/base")
}

func (h AppHandler) HandleGetAddInvestment(ctx *fiber.Ctx) error {

	context := fiber.Map{}
	return ctx.Render("dashboard/add-investment", context, "layout/base")
}

func (h AppHandler) AddInvestment(ctx *fiber.Ctx) error {
	var params data.Investment

	if err := ctx.BodyParser(&params); err != nil {
		return flash.WithError(ctx, fiber.Map{"message": "Invalide body request"}).Redirect("/dashboard/add-investments")
	}

	investment := models.Investment{
		Title:              params.Title,
		Price:              params.Price,
		PercentageIncrease: params.PercentageIncrease,
		ReferralBonus:      params.ReferralBonus,
		Duration:           params.Duration,
	}
	if result := h.DB.Create(&investment); result.Error != nil {
		return flash.WithError(ctx, fiber.Map{"message": "Could not save investment to db"}).Redirect("/dashboard/add-investment")
	}
	return flash.WithSuccess(ctx, fiber.Map{"message": "Successful"}).Redirect("/dashboard/investments")
}

func (h AppHandler) UpdatePicture(ctx *fiber.Ctx) error {

	image, err := ctx.FormFile("image")
	if err != nil {
		return flash.WithError(ctx, fiber.Map{"message": "Invalid Image"}).Redirect("/dashboard/settings")
	}
	err = ctx.SaveFile(image, "./uploads/"+image.Filename)
	if err != nil {
		return flash.WithWarn(ctx, fiber.Map{"message": "An error occured"}).Redirect("/dashboard/settings")
	}

	user, err := h.GetAuthenticatedUser(ctx)
	if err != nil {
		return err
	}
	user.Image = image.Filename
	if result := h.DB.Save(&user); result.Error != nil {
		return flash.WithError(ctx, fiber.Map{"message": "An error occured"}).Redirect("/dashboard/settings")
	}
	return flash.WithSuccess(ctx, fiber.Map{"message": "Updated successfully"}).Redirect("/dashboard/settings")

}
