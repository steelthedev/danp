package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

type AppHandler struct {
	DB    *gorm.DB
	Store *session.Store
}

func (h AppHandler) HandleGetHome(ctx *fiber.Ctx) error {
	context := fiber.Map{}
	return ctx.Render("index", context)
}
