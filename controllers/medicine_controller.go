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

	return c.JSON(fiber.Map{"success": "Success create medicine"})
}

func UpdateMedicine(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("id"), 10, 64)
	medicine := new(models.Medicine)
	var medicineSelected models.Medicine
	if err := c.BodyParser((medicine)); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Request"})
	}

	database.DB.Where("id=?", id).First(&medicineSelected)
	database.DB.Model(&medicine).Updates(models.Medicine{MedicineName: medicineSelected.MedicineName, MedicinePackage: medicineSelected.MedicinePackage, MedicinePrice: medicine.MedicinePrice})
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

// function for get medicine from database
func GetMedicine(c *fiber.Ctx) ([]models.Medicine, error) {
	var medicine []models.Medicine
	result := database.DB.Order("id ASC").Find(&medicine)
	if result.Error != nil {
		return nil, result.Error
	}

	return medicine, nil
}
