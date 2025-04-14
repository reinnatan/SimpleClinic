package routes

import (
	"github.com/gofiber/fiber/v2"
	"reinnatan.com/controllers"
)

func AdminRoutes(app *fiber.App) {
	adminGroup := app.Group("/")
	adminGroup.Get("/section/:name", controllers.SectionAdminMenu)

	//routes for poli
	poliGroup := app.Group("/poli")
	poliGroup.Post("/add", controllers.AddPoli)
	poliGroup.Delete("/delete/:id", controllers.DeletePoli)
	poliGroup.Put("/update/:id", controllers.UpdatePoli)

	//routes for medicine
	medicineGroup := app.Group("/medicine")
	medicineGroup.Post("/add", controllers.AddMedicine)
	medicineGroup.Delete("/delete/:id", controllers.DeleteMedicine)
	medicineGroup.Put("/update/:id", controllers.UpdateMedicine)
}
