package main

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"github.com/steelthedev/danp/handlers"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Info(err.Error())
	}

	viewsEngine := html.New("./public", ".html")

	app := fiber.New(fiber.Config{
		Views:             viewsEngine,
		ViewsLayout:       "layout/main",
		PassLocalsToViews: true,
	})

	app.Static("/static", "./assets")

	appHandler := handlers.AppHandler{}

	// Get Handlers
	app.Get("/", appHandler.HandleGetHome)

	app.Listen(":3000")
}
