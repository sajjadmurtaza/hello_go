package routes

import (
	"fiber_api/db"
	"fiber_api/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	Email     string `json:"email" gorm:"unique"`
}

func CreateUserResponse(userModel *models.User) User {
	return User{
		ID:        userModel.ID,
		FirstName: userModel.FirstName,
		Email:     userModel.Email,
	}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	fmt.Println("Body as string:", string(c.Body()))

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request lol"})
	}

	fmt.Println("Parsed user:", user)

	db.DatabaseRef.Db.Create(&user)
	response := CreateUserResponse(&user)
	return c.Status(201).JSON(response)
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	db.DatabaseRef.Db.Find(&users)

	return c.JSON(users)
}
