package user

import (
	"core/internal/models"
	"core/internal/repositories/auth"
	"core/internal/repositories/user"
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

func (s *Service) GetAllUsers() []models.UserService {
	var usersService []models.UserService

	for _, v := range s.userRepository.GetAllUsers() {
		var userService models.UserService

		userService.ID = v.ID
		userService.CreatedAt = v.CreatedAt
		userService.Login = v.Login
		userService.FirstName = v.FirstName
		userService.LastName = v.LastName
		userService.MiddleName = v.MiddleName
		userService.MainImage = v.MainImage
		userService.SmallImage = v.SmallImage

		usersService = append(usersService, userService)
	}

	return usersService
}

func (s *Service) GetUser(id int64) models.UserService {
	var userService models.UserService
	v := s.userRepository.GetUser(id)

	userService.ID = v.ID
	userService.CreatedAt = v.CreatedAt
	userService.Login = v.Login
	userService.FirstName = v.FirstName
	userService.LastName = v.LastName
	userService.MiddleName = v.MiddleName
	userService.MainImage = v.MainImage
	userService.SmallImage = v.SmallImage

	return userService
}
