package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"reinnatan.com/routes"
)

type User struct {
	Username string
	Password string
}

//ref design
//https://www.figma.com/design/abG5Bvxa9mRuYpUpUwNXff/Design-WEB-Polyclinic?node-id=0-1&p=f

func main() {

	app := fiber.New()
	var store = session.New()

	app.Get("/", func(c *fiber.Ctx) error {
		sess, _ := store.Get(c)
		sess.Destroy()
		return c.SendFile("./public/index.html")
	})

	routes.AdminRoutes(app)

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

		var payload User
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
