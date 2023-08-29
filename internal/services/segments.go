package services

import (
	"github.com/burravlev/avito-tech-test/internal/models"
	"github.com/burravlev/avito-tech-test/internal/repositories"
)

type segemnts struct {
	repository repositories.Segment
}

func SegmentService(repository repositories.Segment) Segment {
	return &segemnts{repository}
}

func (s *segemnts) Save(segemnt *models.Segment) (*models.Segment, error) {
	return nil, nil
}
