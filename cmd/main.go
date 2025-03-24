package main

import (
	"fmt"
	"simple_product_listing_go/internal/routes"
	"simple_product_listing_go/pkg/database"
)
import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	client := database.ConnectDb()

	routes.ProductRoutes(app, client)

	app.Listen(":3000")
	fmt.Println("Server is running on port 3000")
}
