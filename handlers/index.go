package handlers

import "github.com/gofiber/fiber/v2"

type AppHandler struct{}

func (h AppHandler) HandleGetHome(ctx *fiber.Ctx) error {
	context := fiber.Map{}
	return ctx.Render("index", context)
}
