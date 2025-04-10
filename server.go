package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"reinnatan.com/database"
	"reinnatan.com/models"
	"reinnatan.com/request"
	"reinnatan.com/routes"
)

//ref design
//https://www.figma.com/design/abG5Bvxa9mRuYpUpUwNXff/Design-WEB-Polyclinic?node-id=0-1&p=f

func main() {

	app := fiber.New()

	// CORS Middleware (Allows requests from frontend)
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Method() == "OPTIONS" {
			return c.SendStatus(200)
		}
		return c.Next()
	})

	var store = session.New()

	//grouping controller to admin
	routes.AdminRoutes(app)

	//connect database
	database.ConnectDB()

	//Auto migrate models
	//database.DB.Migrator().DropTable(&models.Medicine{})
	database.DB.AutoMigrate(&models.Doctor{})
	database.DB.AutoMigrate(&models.Poli{})
	database.DB.AutoMigrate(&models.Medicine{})

	app.Get("/", func(c *fiber.Ctx) error {
		sess, _ := store.Get(c)
		sess.Destroy()
		return c.SendFile("./public/index.html")
	})

	app.All("/login", func(c *fiber.Ctx) error {
		session, err := store.Get(c)
		if session.Get("user") == "admin" {
			content := "Halaman dashboard"
			title := "Dashboard"
			section := "Dashboard"

			errParse := c.Render("./public/dist/pages/index.html",
				fiber.Map{
					"Content": content,
					"Title":   title,
					"Section": section,
				})
			if errParse != nil {
				return errParse
			}

			c.Set("Content-Type", "text/html")
			return c.Send(c.Response().Body())
		}

		var payload request.User
		if err := c.BodyParser(&payload); err != nil {
			return c.SendFile("./public/index.html")
		}

		if err == nil {
			if payload.Username == "admin" && payload.Password == "admin" {
				session.Set("user", "admin")
				session.Save()

				content := "Halaman dashboard"
				title := "Dashboard"
				section := "Dashboard"

				errParse := c.Render("./public/dist/pages/index.html",
					fiber.Map{
						"Content": content,
						"Title":   title,
						"Section": section,
					})
				if errParse != nil {
					return errParse
				}

				c.Set("Content-Type", "text/html")
				return c.Send(c.Response().Body())
			}
		}

		return c.SendFile("./public/index.html")
	})

	app.Static("/", "./public")

	app.Listen(":3000")
}
