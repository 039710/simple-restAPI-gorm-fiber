package route

import (
	handler "go-learn/handlers"

	"github.com/gofiber/fiber/v2"
)

// SubjectRoute func
func SubjectRoute(app *fiber.App) {

	// routes
	api := app.Group("/subject")

	api.Get("/", handler.GetSubject)
	api.Get("/:id", handler.GetSubjectById)
	api.Post("/", handler.CreateSubject)
	api.Delete("/:id", handler.DeleteSubject)

}
