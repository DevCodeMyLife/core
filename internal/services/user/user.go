package user

import (
	"core/internal/models"
	"core/internal/repositories/auth"
	"core/internal/repositories/user"
	"net/http"
)

type Service struct {
	userRepository *user.Repository
	authRepository *auth.Repository
}

func NewUserService(userRepository *user.Repository, authRepository *auth.Repository) *Service {
	return &Service{
		userRepository: userRepository,
		authRepository: authRepository,
	}
}

func (s *Service) GetAllUsers(token *http.Cookie) []models.User {
	if s.authRepository.IsAuth(token) {
		return s.userRepository.GetAllUsers()
	}
	return nil
}
