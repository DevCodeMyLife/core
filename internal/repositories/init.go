package repositories

import (
	"core/internal/models"
	"core/internal/repositories/auth"
	"core/internal/repositories/user"
	"github.com/ivahaev/go-logger"
	"gorm.io/gorm"
)

type Repositories struct {
	db *gorm.DB

	User *user.Repository
	Auth *auth.Repository
}

func NewRepositories(db *gorm.DB) *Repositories {
	err := db.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		logger.Error(err)
	}

	userRepository := user.NewUserRepository(db)
	authRepository := auth.NewAuthRepository(db)

	return &Repositories{
		db:   db,
		User: userRepository,
		Auth: authRepository,
	}
}
