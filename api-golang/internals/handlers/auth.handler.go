package handlers

import (
	"hms-backend/internals/models"
	"hms-backend/internals/services"

	"github.com/gofiber/fiber/v3"
)

type AuthHandler struct {
	svc *services.Services
}

func (h *AuthHandler) Register(c fiber.Ctx) error {
	user := new(models.User)
	if err := c.Bind().Body(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("request body not provided")
	}

	if err := user.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	user, err := h.svc.Auth.Register(c.Context(), user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	return nil
}
