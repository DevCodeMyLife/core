package models

import (
	"time"
)

type FeedService struct {
	ID uint `json:"id"`

	CreatedAt time.Time `json:"created_at"`

	UID uint `json:"uid"`

	Title string `json:"title"`

	MainImage  string `json:"main_image"`
	SmallImage string `json:"small_image"`

	Value string `json:"value"`

	Login string `json:"login"`

	MainImageProfile  string `json:"main_image_profile"`
	SmallImageProfile string `json:"small_image_profile"`

	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
}
