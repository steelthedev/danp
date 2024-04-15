package handlers

import "github.com/gofiber/fiber/v2"

func (h AppHandler) HanldeGetLogin(ctx *fiber.Ctx) error {
	context := fiber.Map{}
	return ctx.Render("sign-in", context)
}
