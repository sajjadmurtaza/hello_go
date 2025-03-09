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

	fmt.Println("Empty user struct:", user)
	fmt.Println("Raw body bytes:", c.Body())
	fmt.Println("Body as string:", string(c.Body()))
	fmt.Println("BodyParser result:", c.BodyParser(&user))

	if err := c.BodyParser(&user); err != nil {
		fmt.Println("Parse error:", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request lol"})
	}

	fmt.Println("Parsed user:", user)

	db.DatabaseRef.Db.Create(&user)

	response := CreateUserResponse(&user)

	return c.Status(201).JSON(response)
}
