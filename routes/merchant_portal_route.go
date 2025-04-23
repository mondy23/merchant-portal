package routes

import (
	"merchant-portal/pkg/features/access/handlers"
	"merchant-portal/pkg/features/access/repositories"
	"merchant-portal/pkg/features/access/services"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	db *pgxpool.Pool
)

func MerchantPortalRoute(app *fiber.App) {
	repo := repositories.NewAccessRepository(db)
	services := services.NewAccessServices(repo)
	handlers := handlers.NewAccessHandlers(services)

	// Access/Portal User Routes
	app.Get("/portal-users", handlers.GetPortalUsers)          // Get all portal users (list)
	app.Post("/portal-users", handlers.AddPortalUser)          // Create a new portal user
	app.Get("/portal-users/:id", handlers.GetPortalUser)       // Get a specific portal user by ID
	app.Put("/portal-users/:id", handlers.UpdatePortalUser)    // Update an existing portal user by ID
	app.Delete("/portal-users/:id", handlers.DeletePortalUser) // Delete a portal user by ID
}
