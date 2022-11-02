package main

import (
	"go-w-golang/basic/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/public", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Google Dev!")
	})

	// Basic routes for GET, POST, PUT, PATCH, DELETE
	newRoute := app.Group("/new-route")

	// newRoute.Get("/", handlers.Get)
	// newRoute.Post("/", handlers.Post)
	// newRoute.Put("/", handlers.Put)
	// newRoute.Patch("/", handlers.Patch)
	// newRoute.Delete("/", handlers.Delete)

	newRoute.Get("/*", handlers.Info)
	newRoute.Get("/", handlers.Info)
	newRoute.Post("/", handlers.Info)
	newRoute.Put("/", handlers.Info)
	newRoute.Patch("/", handlers.Info)
	newRoute.Delete("/", handlers.Info)

	// Parameterized routes
	app.Get("/test-params/*", handlers.WildCard)
	app.Get("resturant/:id/food/:food_id", handlers.Resturant)

	// Books Group
	books := app.Group("books")

	books.Get("/", handlers.GetBooks)
	books.Get("/:id", handlers.GetBook)
	books.Post("/", handlers.CreateBook)
	books.Delete("/:id", handlers.DeleteBook)

	// Download file
	app.Get("/download", handlers.Download)

	app.Listen(":3000")
}
