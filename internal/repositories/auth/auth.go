package auth

import (
	"core/internal/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) IsAuth(token string) bool {
	var u []models.User
	r.db.Table("users").Where("token = ?", token).Find(&u)

	if len(u) == 1 {
		return true
	}

	return false
}
