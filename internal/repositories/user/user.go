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
	r.db.Table("users").Where("deleted_at is null").Find(&u)

	return u
}

func (r *Repository) GetUserByID(id int64) models.User {
	var u models.User
	r.db.Table("users").Where("id = ? AND deleted_at is null", id).Find(&u)

	return u
}
