package controllers

import (
	"context"
	"strings"

	"github.com/criticic/cupid-backend/random_name_generator"
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

	email := token.Claims["email"].(string)
	// user_firestore, err := utils.FirestoreClient.Collection("users").Doc(token.UID).Get(context.Background())
	_, err = utils.FirestoreClient.Collection("users").Doc(token.UID).Get(context.Background())
	if err != nil {
		if !isAllowedEmail(email) {
			err := utils.AuthClient.DeleteUser(context.Background(), token.UID)
			if err != nil {
				log.Errorf("error deleting user: %v\n", err)
			}

			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Email should end with @itbhu.ac.in",
			})

		} else {
			username := generateUniqueUsername()
			log.Info("Creating user in Firestore: ", username)
			utils.FirestoreClient.Collection("users").Doc(token.UID).Set(context.Background(), map[string]interface{}{
				"email":    email,
				"username": username,
			})
		}
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

func generateUniqueUsername() string {
	username := random_name_generator.GenerateName()
	for usernameExists(username) {
		username = random_name_generator.GenerateName()
	}
	return username
}

func usernameExists(username string) bool {
	iter := utils.FirestoreClient.Collection("users").Where("username", "==", username).Documents(context.Background())
	// Check if the username exists
	if _, err := iter.Next(); err != nil {
		return false
	}
	return true
}
