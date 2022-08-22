package services

import (
	"core/internal/repositories"
	"core/internal/services/auth"
	"core/internal/services/feed"
	"core/internal/services/user"
)

type Services struct {
	repo *repositories.Repositories

	User *user.Service
	Auth *auth.Service
	Feed *feed.Service
}

func NewServices(repo *repositories.Repositories) *Services {
	userService := user.NewUserService(repo.User)
	authService := auth.NewAuthService(repo.Auth)
	feedService := feed.NewFeedService(repo.Feed)

	return &Services{
		repo: repo,
		User: userService,
		Auth: authService,
		Feed: feedService,
	}
}
