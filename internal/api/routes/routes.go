package routes

import (
	"github.com/backsoul/doberman/internal/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/send", controllers.Hello)
}
