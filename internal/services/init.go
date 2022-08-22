package services

import (
	"core/internal/repositories"
	"core/internal/services/auth"
	"core/internal/services/user"
)

type Services struct {
	repo *repositories.Repositories

	User *user.Service
	Auth *auth.Service
}

func NewServices(repo *repositories.Repositories) *Services {
	userService := user.NewUserService(repo.User)
	authService := auth.NewAuthService(repo.Auth)

	return &Services{
		repo: repo,
		User: userService,
		Auth: authService,
	}
}
