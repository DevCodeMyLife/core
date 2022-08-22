package auth

import (
	"core/internal/repositories/auth"
)

type Service struct {
	authRepository *auth.Repository
}

func NewAuthService(authRepository *auth.Repository) *Service {
	return &Service{
		authRepository: authRepository,
	}
}

func (s *Service) IsAuth(token string) bool {
	return s.authRepository.IsAuth(token)
}
