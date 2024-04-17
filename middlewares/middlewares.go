package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/steelthedev/danp/handlers"
	"github.com/sujit-baniya/flash"
)

type MiddleWareHandler struct {
	*handlers.AppHandler
}

func (h MiddleWareHandler) MustBeAuthenticated(ctx *fiber.Ctx) error {
	// Check if AppHandler is initialized
	if h.AppHandler == nil {
		return fiber.NewError(fiber.StatusInternalServerError, "AppHandler is nil")
	}

	// Retrieve the authenticated user
	user, err := h.GetAuthenticatedUser(ctx)
	if err != nil {
		return err
	}

	// Check if user is authenticated
	if user == nil {
		// Redirect to home page or login page if user is not authenticated
		return ctx.RedirectBack("/")
	}

	// Continue to the next middleware or handler
	return ctx.Next()
}

func WithFlash(ctx *fiber.Ctx) error {
	values := flash.Get(ctx)
	ctx.Locals("flash", values)
	return ctx.Next()
}
