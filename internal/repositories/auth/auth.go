package auth

import (
	"gorm.io/gorm"
	"net/http"
)

type Repository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) IsAuth(token *http.Cookie) bool {
	var t = struct {
		token string
	}{}

	r.db.Table("users").Where("token = ?", token.Value).Find(&t)

	if t.token != "" {
		return true
	}

	return false
}
