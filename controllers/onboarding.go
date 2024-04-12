package controllers

import (
	"context"
	"log"

	"github.com/criticic/cupid-backend/utils"
	"github.com/gofiber/fiber/v3"
)

func GetUsername(c fiber.Ctx) error {
	idToken := c.Query("token")

	token, err := utils.AuthClient.VerifyIDTokenAndCheckRevoked(context.Background(), idToken)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error verifying token",
		})
	}

	user, err := utils.FirestoreClient.Collection("users").Doc(token.UID).Get(context.Background())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	log.Print("User found: ", user.Data()["username"])

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"username": user.Data()["username"],
	})
}
