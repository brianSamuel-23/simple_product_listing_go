package routes

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"simple_product_listing_go/internal/handler"
	"simple_product_listing_go/internal/middleware"
	"simple_product_listing_go/internal/repository"
	"simple_product_listing_go/internal/service"
)

func ProductRoutes(a *fiber.App, client *mongo.Client) {

	productRepo := repository.NewProductRepository(client)
	productPriceRepo := repository.NewProductPriceRepository(client)
	productService := service.NewProductService(productRepo, productPriceRepo)
	productHandler := handler.NewProductHandler(productService)

	route := a.Group("v1/product", middleware.JwtMiddleware())
	route.Get("/", productHandler.GetProducts)
	route.Get("/prices", productHandler.GetProductPrices)
}
