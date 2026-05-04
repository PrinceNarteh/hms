package routes

import (
	"hms-backend/internals/handlers"

	"github.com/gofiber/fiber/v3"
)

func authRoutes(app fiber.Router, h *handlers.Handlers) {
	r := app.Group("/auth")

	r.Post("/login", h.Auth.Login)
	r.Post("/register", h.Auth.Register)
}
