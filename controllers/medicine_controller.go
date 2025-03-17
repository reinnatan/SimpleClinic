package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"reinnatan.com/database"
	"reinnatan.com/models"
)

func AddMedicine(c *fiber.Ctx) error {
	medicine := new(models.Medicine)

	if err := c.BodyParser(medicine); err != nil {
		return c.Status(503).JSON(fiber.Map{"error": "Couldn't create medicine"})
	}

	result := database.DB.Create(medicine)
	if result.Error != nil {
		return c.Status(503).JSON(fiber.Map{"error": "Couldn't create medicine"})
	}

	return c.JSON(medicine)
}

func UpdateMedicine(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("id"), 10, 64)
	medicine := new(models.Medicine)
	var medicineSelected models.Medicine
	if err := c.BodyParser((medicine)); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Request"})
	}

	database.DB.Where("id=?", id).First(&medicineSelected)
	database.DB.Model(&medicine).Updates(models.Medicine{Name: medicineSelected.Name, Package: medicineSelected.Package, Price: medicine.Price})
	return c.JSON(fiber.Map{"message": "successfully updated"})
}

func DeleteMedicine(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("id"), 10, 64)
	result := database.DB.Delete(models.Poli{ID: id})
	if result.Error != nil {
		return c.Status(500).SendString("Failed to delete user")
	}
	return c.JSON(fiber.Map{"message": "User deleted successfully"})
}
