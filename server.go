package main

import (
	"bytes"
	"html/template"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Username string
	Password string
}

//ref design
//https://www.figma.com/design/abG5Bvxa9mRuYpUpUwNXff/Design-WEB-Polyclinic?node-id=0-1&p=f

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		//content := ""
		//return c.Render("./public/dist/pages/index.html", fiber.Map{"Content": content})
		return c.SendFile("./public/index.html")
	})

	app.Get("/section/:name", func(c *fiber.Ctx) error {

		sectionName := c.Params("name")
		content := ""
		title := ""
		section := ""

		switch sectionName {
		case "dashboard":
			content = "Halaman dashboard"
			title = "Dashboard"
			section = "Dashboard"

		case "dokter":
			tmpl, err := template.ParseFiles("./public/dist/pages/dokter.html")
			if err != nil {
				return err
			}

			//parsing data to template
			data := map[string]string{}

			var buf bytes.Buffer
			err = tmpl.Execute(&buf, data)
			if err != nil {
				return err
			}
			title = "Tambah / Edit Dokter"
			section = "Dokter"
			renderedHTML := buf.String()
			content = string(renderedHTML)

		case "pasien":
			tmpl, err := template.ParseFiles("./public/dist/pages/pasien.html")
			if err != nil {
				return err
			}

			//parsing data to template
			data := map[string]string{}

			var buf bytes.Buffer
			err = tmpl.Execute(&buf, data)
			if err != nil {
				return err
			}
			title = "Tambah / Edit Pasien"
			section = "Pasien"
			renderedHTML := buf.String()
			content = string(renderedHTML)

		case "poli":
			tmpl, err := template.ParseFiles("./public/dist/pages/poli.html")
			if err != nil {
				return err
			}

			//parsing data to template
			data := map[string]string{}

			var buf bytes.Buffer
			err = tmpl.Execute(&buf, data)
			if err != nil {
				return err
			}
			title = "Tambah / Edit Poli"
			section = "Poli"
			renderedHTML := buf.String()
			content = string(renderedHTML)

		case "obat":
			tmpl, err := template.ParseFiles("./public/dist/pages/obat.html")
			if err != nil {
				return err
			}

			//parsing data to template
			data := map[string]string{}

			var buf bytes.Buffer
			err = tmpl.Execute(&buf, data)
			if err != nil {
				return err
			}
			title = "Tambah / Edit Obat"
			section = "Obat"
			renderedHTML := buf.String()
			content = string(renderedHTML)

		case "signout":
			errSignout := c.Render("./public/index.html",
				fiber.Map{
					"Content": content,
					"Title":   title,
					"Section": section,
				})
			if errSignout != nil {
				return errSignout
			}
			c.Set("Content-Type", "text/html")
			return c.Send(c.Response().Body())

		default:
			break
		}

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
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		payload := new(User)
		if err := c.BodyParser(payload); err != nil {
			return err
		}

		if payload.Username == "admin" && payload.Password == "admin" {
			return c.SendFile("./public/dist/pages/index.html")
		}
		return c.SendFile("./public/index.html")
	})

	app.Static("/", "./public")

	app.Listen(":3000")
}
