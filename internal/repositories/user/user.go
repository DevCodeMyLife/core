package user

import (
	"core/internal/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetAllUsers() []models.User {
	var u []models.User
	r.db.Table("users").Find(&u)

	return u
}

func (r *Repository) GetUser(id int64) models.User {
	var u models.User
	r.db.Table("users").Where("id = ?", id).Find(&u)

	return u
}
