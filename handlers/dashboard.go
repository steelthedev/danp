package handlers

import (
	"github.com/gofiber/fiber/v2"
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
