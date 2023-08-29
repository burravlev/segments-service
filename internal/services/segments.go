package services

import (
	"errors"
	"fmt"

	"github.com/burravlev/avito-tech-test/internal/models"
	"github.com/burravlev/avito-tech-test/internal/repositories"
	"gorm.io/gorm"
)

type segemnts struct {
	repository repositories.Segment
}

func SegmentService(repository repositories.Segment) Segment {
	return &segemnts{repository}
}

func (s *segemnts) Save(segemnt *models.Segment) error {
	err := s.repository.Save(segemnt)
	if errors.Is(gorm.ErrDuplicatedKey, err) {
		return fmt.Errorf("duplicated key")
	}
	return err
}
