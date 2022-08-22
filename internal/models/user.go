package models

import "time"

type UserService struct {
	ID uint `json:"id"`

	CreatedAt time.Time `json:"created_at"`

	Login string `json:"login"`

	MainImage  string `json:"main_image"`
	SmallImage string `json:"small_image"`

	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
}
