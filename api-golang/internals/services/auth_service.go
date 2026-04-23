package services

import (
	"context"
	"errors"

	"hms-backend/internals/models"
	"hms-backend/internals/repository"
)

var _ AuthService = (*authService)(nil)

type AuthService interface {
	Register(context.Context, *models.User) (*models.User, error)
}

type authService struct {
	repo *repository.Repository
}

func (s *authService) Register(ctx context.Context, user *models.User) (*models.User, error) {
	userExists, err := s.repo.User.FindOneByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	if userExists != nil {
		return nil, errors.New(`user with email "%s" already exists`)
	}

	if err := user.HashPassword(); err != nil {
		return nil, err
	}

	if err := s.repo.User.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) Login(ctx context.Context, data *models.LoginDTO) (*models.User, error) {
	user, err := s.repo.User.FindOneByEmail(ctx, data.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !user.CompareHashandPassword(data.Password) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
