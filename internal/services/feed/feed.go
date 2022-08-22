package feed

import (
	"core/internal/models"
	"core/internal/repositories/feed"
)

type Service struct {
	feedRepository *feed.Repository
}

func NewFeedService(feedRepository *feed.Repository) *Service {
	return &Service{
		feedRepository: feedRepository,
	}
}

func (s *Service) GetAllFeeds() []models.FeedService {
	return s.feedRepository.GetAllFeeds()
}

func (s *Service) GetFeedByID(id int64) models.FeedService {
	return s.feedRepository.GetFeedByID(id)
}
