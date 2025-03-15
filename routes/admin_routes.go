package routes

import (
	"github.com/gofiber/fiber/v2"
	"reinnatan.com/controllers"
)

func AdminRoutes(app *fiber.App) {
	adminGroup := app.Group("/")
	adminGroup.Get("/section/:name", controllers.SectionAdminMenu)
}
