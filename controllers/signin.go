package controllers

import (
	"context"
	"strings"

	"github.com/criticic/cupid-backend/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func SignIn(c fiber.Ctx) error {
	// Get the token from query
	idToken := c.Query("token")
	log.Info("idToken: ", idToken)

	// Verify the token
	token, err := utils.AuthClient.VerifyIDTokenAndCheckRevoked(context.Background(), idToken)
	if err != nil {
		log.Errorf("error verifying ID token: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error verifying token",
		})
	}

	// Make sure email ends with itbhu.ac.in
	email := token.Claims["email"].(string)
	if !isAllowedEmail(email) {
		err := utils.AuthClient.DeleteUser(context.Background(), token.UID)
		if err != nil {
			log.Errorf("error deleting user: %v\n", err)
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email should end with @itbhu.ac.in",
		})

	} else {
		utils.FirestoreClient.Collection("users").Doc(token.UID).Set(context.Background(), map[string]interface{}{
			"email": email,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User signed in successfully",
	})
}

var allowedDomains = []string{
	"itbhu.ac.in",
	"iitbhu.ac.in",
}

func isAllowedEmail(email string) bool {
	for _, domain := range allowedDomains {
		if strings.HasSuffix(email, "@"+domain) {
			return true
		}
	}
	return false
}
