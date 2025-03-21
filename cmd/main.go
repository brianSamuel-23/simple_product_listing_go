package main

import "fmt"
import "github.com/gofiber/fiber/v2"


func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
	fmt.Println("Server is running on port 3000")
}