package services

import "github.com/burravlev/avito-tech-test/internal/models"

type Segment interface {
	// Save saves or updated segment
	Save(*models.Segment) (*models.Segment, error)
}
