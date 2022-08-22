package feed

import (
	"core/internal/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewFeedRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetAllFeeds() []models.FeedService {
	var f []models.FeedService
	r.db.Table("feeds f").Select(
		"f.id, f.created_at, f.uid, f.title, f.main_image, f.small_image, f.value, e.login, e.main_image as main_image_profile, e.small_image as small_image_profile, e.first_name, e.middle_name, e.last_name",
	).Joins(
		"inner join users e on f.uid = e.id",
	).Where(
		"f.deleted_at is null AND e.deleted_at is null",
	).Find(&f)

	return f
}

func (r *Repository) GetFeedByID(id int64) models.FeedService {
	var f models.FeedService
	r.db.Table("feeds f").Select(
		"f.id, f.created_at, f.uid, f.title, f.main_image, f.small_image, f.value, e.login, e.main_image as main_image_profile, e.small_image as small_image_profile, e.first_name, e.middle_name, e.last_name",
	).Joins(
		"inner join users e on f.uid = e.id",
	).Where("f.id = ? AND f.deleted_at is null AND e.deleted_at is null", id).Find(&f)

	return f
}
