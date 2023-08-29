package services

import (
	"errors"
	"fmt"

	"github.com/burravlev/avito-tech-test/internal/models"
	"github.com/burravlev/avito-tech-test/internal/repositories"
	"gorm.io/gorm"
)

type segments struct {
	repository repositories.Segment
}

func SegmentService(repository repositories.Segment) Segment {
	return &segments{repository}
}

func (s *segments) Save(segemnt *models.Segment) error {
	err := s.repository.Save(segemnt)
	if errors.Is(gorm.ErrDuplicatedKey, err) {
		return fmt.Errorf("duplicated key")
	}
	return err
}

func (s *segments) Delete(name string) error {
	return s.repository.Delete(name)
}

func (s *segments) Add(userId uint, segments []string) {

}

func (s *segments) GetByUser(userId uint) ([]models.Segment, error) {
	return s.repository.GetByUser(userId)
}
