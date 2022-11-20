package handler

import (
	"go-learn/db"
	model "go-learn/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

// create handler get all, get by id, create, update, delete
func Login(c *fiber.Ctx) error {
	db := db.DB
	var student model.Student
	c.BodyParser(&student)
	db.Where("email = ? AND password = ?", student.Email, jwt.EncodeSegment([]byte(student.Password))).First(&student)
	// if student is not found
	if student.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}
	// create jwt token
	token := jwt.New(jwt.SigningMethodHS256)
	// set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = student.ID
	claims["name"] = student.Name
	claims["email"] = student.Email
	// generate encoded token and send it as response
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error on login",
		})
	}
	return c.JSON(fiber.Map{
		"token": t,
	})
}

func GetStudent(c *fiber.Ctx) error {
	db := db.DB
	var student []model.Student
	// nested preload
	db.Preload("Subjects.Students", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "email", "name", "dob", "students")
	}).Find(&student)

	return c.JSON(student)
}

func GetStudentById(c *fiber.Ctx) error {
	db := db.DB
	id, _ := strconv.Atoi(c.Params("id"))
	var student model.Student
	db.Find(&student, id)
	if student.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Student not found",
		})
	}
	return c.JSON(student)
}

func CreateStudent(c *fiber.Ctx) error {
	db := db.DB
	var student model.Student
	c.BodyParser(&student)
	db.Create(&student)
	return c.JSON(student)
}

func UpdateStudent(c *fiber.Ctx) error {
	db := db.DB
	id, _ := strconv.Atoi(c.Params("id"))
	var student model.Student
	db.Find(&student, id)
	c.BodyParser(&student)
	// if student is not found
	if student.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Student not found",
		})
	}
	// if password in request is not empty
	if student.Password != "" {
		student.Password = jwt.EncodeSegment([]byte(student.Password))
	}

	db.Save(&student)
	return c.JSON(student)
}

func DeleteStudent(c *fiber.Ctx) error {
	db := db.DB
	id, _ := strconv.Atoi(c.Params("id"))
	var student model.Student
	db.Find(&student, id)
	// if student is not found
	if student.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Student not found",
		})
	}
	db.Delete(&student)
	return c.JSON(fiber.Map{
		"message": "Student deleted",
	})
}
