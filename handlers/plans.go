package handlers

import "github.com/gofiber/fiber/v2"

func (h AppHandler) HandleGetPlan(ctx *fiber.Ctx) error {
	context := fiber.Map{}
	return ctx.Render("plan", context)
}
