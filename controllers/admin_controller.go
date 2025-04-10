package controllers

import (
	"bytes"
	"html/template"

	"github.com/gofiber/fiber/v2"
)

func SectionAdminMenu(c *fiber.Ctx) error {
	sectionName := c.Params("name")
	content := "Halaman dashboard"
	title := "Dashboard"
	section := "Dashboard"

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
		polis, err := GetPoli(c)
		if err != nil {
			return err
		}
		data := fiber.Map{
			"Polis": polis,
		}

		var buf bytes.Buffer
		err = tmpl.Execute(&buf, data)
		if err != nil {
			return err
		}
		title = "Tambah / Edit Poli"
		section = "Poli"
		renderedHTML := buf.String()
		content = string(renderedHTML)

	case "medicine":
		tmpl, err := template.ParseFiles("./public/dist/pages/medicine.html")
		if err != nil {
			return err
		}

		medicines, err := GetMedicine(c)
		if err != nil {
			return err
		}

		//parsing data to template
		data := fiber.Map{
			"Medicines": medicines,
		}

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
}

func AddDoctor(c *fiber.Ctx) error {
	return nil
}
