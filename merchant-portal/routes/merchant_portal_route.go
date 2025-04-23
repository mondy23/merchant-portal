package routes

import (
	"merchant-portal/pkg/features/access/users/handlers"

	"github.com/gofiber/fiber/v2"
)

func MerchantPortalRoute(app *fiber.App) {
	// Access/Portal User Routes
	app.Get("/portal-users", handlers.GetPortalUsers)          // Get all portal users (list)
	app.Post("/portal-users", handlers.AddPortalUser)          // Create a new portal user
	app.Get("/portal-users/:id", handlers.GetPortalUser)       // Get a specific portal user by ID
	app.Put("/portal-users/:id", handlers.UpdatePortalUser)    // Update an existing portal user by ID
	app.Delete("/portal-users/:id", handlers.DeletePortalUser) // Delete a portal user by ID
}
