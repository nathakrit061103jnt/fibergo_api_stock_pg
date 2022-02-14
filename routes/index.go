package routes

import (
	"fibergo_api_stock_pg/configs"
	"fibergo_api_stock_pg/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/*", configs.ConfigAuth)
	app.Static("/", "./images")

	// group  route
	api := app.Group("/api/v1")
	product := api.Group("/products", configs.ConfigAuth)
	auth := app.Group("/auth") // auth

	// Product route
	product.Get("/", controllers.GetProductsAll)
	product.Get("/:id", controllers.GetProduct)
	product.Post("/", controllers.CreateProduct)
	product.Put("/:id", controllers.UpdateProduct)
	product.Delete("/:id", controllers.DeleteProduct)

	// auth route
	auth.Post("/login", controllers.Login)
	auth.Post("/register", controllers.Register)

}
