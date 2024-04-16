package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AppHandler struct {
	DB *gorm.DB
}

func (h AppHandler) HandleGetHome(ctx *fiber.Ctx) error {
	context := fiber.Map{}
	return ctx.Render("index", context)
}
