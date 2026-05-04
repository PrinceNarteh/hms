// Package routes
package routes

import (
	"hms-backend/internals/handlers"

	"github.com/gofiber/fiber/v3"
)

func InitRoutes(app fiber.Router, handlers *handlers.Handlers) {
	authRoutes(app, handlers)
}
