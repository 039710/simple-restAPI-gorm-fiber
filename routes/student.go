package route

import (
	handler "go-learn/handlers"
	"go-learn/middlewares"

	"github.com/gofiber/fiber/v2"
)

// studentRoute func
func StudentRoute(app *fiber.App) {

	// routes
	app.Post("/login", handler.Login)

	api := app.Group("/student")
	// use middleware
	api.Use(middlewares.Auth())
	api.Get("/", handler.GetStudent)
	api.Get("/:id", handler.GetStudentById)
	api.Post("/", handler.CreateStudent)
	api.Put("/:id", handler.UpdateStudent)
	api.Delete("/:id", handler.DeleteStudent)

}
