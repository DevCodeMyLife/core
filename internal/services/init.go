package services

import (
	"core/internal/repositories"
	"core/internal/services/user"
)

type Services struct {
	repo *repositories.Repositories

	User *user.Service
}

func NewServices(repo *repositories.Repositories) *Services {
	userService := user.NewUserService(repo.User, repo.Auth)

	return &Services{
		repo: repo,
		User: userService,
	}
}
