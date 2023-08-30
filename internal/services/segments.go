package services

import (
	"github.com/burravlev/avito-tech-test/internal/models"
	"github.com/burravlev/avito-tech-test/internal/repositories"
)

type segment struct {
	repository repositories.Segment
}

func SegmentService(repository repositories.Segment) Segment {
	return &segment{repository}
}

func (s *segment) Create(segment *models.Segment) error {
	return s.repository.Create(segment)
}

func (s *segment) Delete(name string) error {
	return s.repository.Delete(name)
}

func (s *segment) GetUserSegments(userId uint) (*models.User, error) {
	segments, err := s.repository.GetByUserID(userId)
	return &models.User{ID: userId, Segments: segments}, err
}

func (s *segment) UpdateSegments(userId uint, add, delete []string) (*models.User, error) {
	segments, err := s.repository.Update(userId, add, delete)
	return &models.User{ID: userId, Segments: segments}, err
}
