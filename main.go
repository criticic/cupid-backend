package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"

	"github.com/criticic/cupid-backend/controllers"
	"github.com/criticic/cupid-backend/utils"
)

func main() {
	// Create a new router & API.
	app := fiber.New()

	// Initialize Firebase clients.
	utils.InitializeFirebaseClients()

	// Define a route.
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Welcome to Cupid Backend 👋!")
	})

	auth := app.Group("/auth")
	auth.Get("/signin", controllers.SignIn)

	// Start the server on port 3000.
	log.Fatal(app.Listen(":3000"))
}
