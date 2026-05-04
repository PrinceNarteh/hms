// Package handlers
package handlers

import "hms-backend/internals/services"

type Handlers struct {
	Auth *AuthHandler
}

func InitHandlers(svc *services.Services) *Handlers {
	return &Handlers{
		Auth: &AuthHandler{svc: svc},
	}
}
