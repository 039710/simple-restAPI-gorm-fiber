package route

import (
	"github.com/gofiber/fiber/v2"
)

// studentRoute func
func InitRoutes(app *fiber.App) {
	StudentRoute(app)
	SubjectRoute(app)
	// routes
}
