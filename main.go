package main

import (
	"context"

	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"google.golang.org/api/option"
)

func main() {
	// Create a new router & API.
	app := fiber.New()

	opt := option.WithCredentialsFile("serviceAccountKey.json")
	firebase_app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	println("Firebase app initialized", firebase_app)

	auth_client, err := firebase_app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	firestore_client, err := firebase_app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("error getting Firestore client: %v\n", err)
	}

	// Define a route.
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Welcome to Cupid Backend 👋!")
	})

	auth := app.Group("/auth")
	auth.Get("/verifyIdToken", func(c fiber.Ctx) error {

		// Get the token from query
		idToken := c.Query("token")
		log.Info("idToken: ", idToken)

		// Verify the token
		token, err := auth_client.VerifyIDTokenAndCheckRevoked(context.Background(), idToken)
		if err != nil {
			log.Fatalf("error verifying ID token: %v\n", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Error verifying token",
			})
		}

		// Make sure email ends with itbhu.ac.in
		if token.Claims["email"] != nil {
			email := token.Claims["email"].(string)
			if email[len(email)-12:] != "@itbhu.ac.in" {
				// Delete the user
				err := auth_client.DeleteUser(context.Background(), token.UID)
				if err != nil {
					log.Fatalf("error deleting user: %v\n", err)
				}

				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "Email should end with @itbhu.ac.in",
				})
			} else {
				firestore_client.Collection("users").Doc(token.UID).Set(context.Background(), map[string]interface{}{
					"email": email,
				})
			}
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"uid": token.UID,
		})
	})

	// Start the server on port 3000.
	log.Fatal(app.Listen(":3000"))
}
