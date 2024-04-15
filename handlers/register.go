package handlers

import "github.com/gofiber/fiber/v2"

func (h AppHandler) HanldeGetRegister(ctx *fiber.Ctx) error {
	context := fiber.Map{}
	return ctx.Render("sign-up", context)
}
