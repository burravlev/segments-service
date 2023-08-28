package repositories

import "github.com/burravlev/avito-tech-test/internal/models"

type segment struct {
}

func (s *segment) Save(*models.Segment) (*models.Segment, error) {
	return nil, nil
}
