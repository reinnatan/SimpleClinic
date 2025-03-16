package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"reinnatan.com/database"
	"reinnatan.com/models"
)

func AddPoli(c *fiber.Ctx) error {
	poli := new(models.Poli)

	if err := c.BodyParser(poli); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Request"})
	}

	result := database.DB.Create(poli)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Couldn't create user"})
	}

	return c.JSON(poli)
}

func UpdatePoli(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("id"), 10, 64)
	poli := new(models.Poli)
	var poliSelected models.Poli
	if err := c.BodyParser((poli)); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Request"})
	}

	database.DB.Where("id=?", id).First(&poliSelected)
	database.DB.Model(&poliSelected).Updates(models.Poli{PoliName: poli.PoliName, PoliDescription: poli.PoliDescription})
	return c.JSON(fiber.Map{"message": "successfully updated"})
}

func DeletePoli(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("id"), 10, 64)
	result := database.DB.Delete(models.Poli{ID: id})
	if result.Error != nil {
		return c.Status(500).SendString("Failed to delete user")
	}
	return c.JSON(fiber.Map{"message": "User deleted successfully"})
}

// function for get poli from database
func GetPoli(c *fiber.Ctx) ([]models.Poli, error) {
	var polis []models.Poli
	result := database.DB.Order("id ASC").Find(&polis)
	if result.Error != nil {
		return nil, result.Error
	}

	return polis, nil
}
