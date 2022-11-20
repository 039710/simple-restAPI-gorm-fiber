package handler

import (
	"go-learn/db"
	model "go-learn/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// create handler get all, get by id, create, update, delete
func GetSubject(c *fiber.Ctx) error {
	db := db.DB
	var subject []model.Subject
	db.Preload("Students").Find(&subject)
	return c.JSON(subject)
}

func GetSubjectById(c *fiber.Ctx) error {
	db := db.DB
	id, _ := strconv.Atoi(c.Params("id"))
	var subject model.Subject
	db.Preload("Students").Find(&subject, id)
	if subject.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Subject not found",
		})
	}
	return c.JSON(subject)
}

func CreateSubject(c *fiber.Ctx) error {
	db := db.DB
	var subject model.Subject
	c.BodyParser(&subject)
	db.Create(&subject)
	return c.JSON(subject)
}

func UpdateSubject(c *fiber.Ctx) error {
	db := db.DB
	id, _ := strconv.Atoi(c.Params("id"))
	var subject model.Subject
	db.Find(&subject, id)
	if subject.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Subject not found",
		})
	}
	c.BodyParser(&subject)
	db.Save(&subject)
	return c.JSON(subject)
}

func DeleteSubject(c *fiber.Ctx) error {
	db := db.DB
	id, _ := strconv.Atoi(c.Params("id"))
	var subject model.Subject
	db.Find(&subject, id)
	// if subject is not found
	if subject.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Subject not found",
		})
	}
	db.Delete(&subject)
	return c.JSON(fiber.Map{
		"message": "Subject deleted",
	})
}
