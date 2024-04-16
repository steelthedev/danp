package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/steelthedev/danp/handlers"
)

type MiddleWareHandler struct {
	*handlers.AppHandler
}

func (h MiddleWareHandler) MustBeAuthenticated(ctx *fiber.Ctx) error {
	user, err := h.GetAuthenticatedUser(ctx)
	if err != nil {
		return err
	}
	if user == nil {
		return ctx.RedirectBack("/")
	}

	return ctx.Next()
}
