// Package services
package services

import "hms-backend/internals/repository"

type Services struct {
	Auth AuthService
}

func InitServices(repo *repository.Repository) *Services {
	return &Services{
		Auth: &authService{repo: repo},
	}
}
