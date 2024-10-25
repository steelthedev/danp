package main

import (
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres/v3"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"github.com/steelthedev/danp/db"
	"github.com/steelthedev/danp/handlers"
	"github.com/steelthedev/danp/middlewares"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Info(err.Error())
	}
	dbUrl := string(os.Getenv("DB_URL"))
	db := db.Init(dbUrl)

	viewsEngine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:             viewsEngine,
		ViewsLayout:       "layout/main",
		PassLocalsToViews: true,
	})

	app.Static("/static", "./assets")
	app.Static("/uploads", "./uploads")
	// handle sessions
	storage := postgres.New(
		postgres.Config{
			ConnectionURI: os.Getenv("DB_URL"),
			Table:         "fiber",
			Reset:         false,
		},
	)
	store := session.New(session.Config{
		Storage: storage,
	})

	appHandler := handlers.AppHandler{
		DB:    db,
		Store: store,
	}

	middlewareHandler := middlewares.MiddleWareHandler{
		AppHandler: &appHandler,
	}

	app.Use(middlewares.WithFlash)
	// Get Handlers
	app.Get("/", appHandler.HandleGetHome)
	app.Get("/about", appHandler.HandleGetAbout)
	app.Get("/plans", appHandler.HandleGetPlan)
	app.Get("/contact", appHandler.HandleGetContact)
	app.Get("/login", appHandler.HanldeGetLogin)
	app.Get("/register", appHandler.HanldeGetRegister)
	app.Get("/logout", appHandler.Logout)

	//post handlers
	app.Post("/register", appHandler.RegisterUser)
	app.Post("/login", appHandler.LoginUser)

	//dashboard
	dashboard := app.Group("/dashboard", middlewareHandler.MustBeAuthenticated)
	dashboard.Get("/", appHandler.HandleGetDashboard)
	dashboard.Get("/settings", appHandler.HandleGetSettings)
	dashboard.Post("/settings/update-profile", appHandler.EditUser)
	dashboard.Get("/wallet", appHandler.HandleGetWallet)
	dashboard.Get("/investments", appHandler.HandleGetInvestment)
	dashboard.Get("/add-investments", appHandler.HandleGetAddInvestment)
	dashboard.Post("/add-investment", appHandler.AddInvestment)
	dashboard.Post("/upload-profile-picture", appHandler.UpdatePicture)

	app.Listen(":3000")
}
