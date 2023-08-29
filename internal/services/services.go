package services

import "github.com/burravlev/avito-tech-test/internal/models"

type Segment interface {
	// Save saves or updated segment
	Save(*models.Segment) error
	// Delete segment by slug name
	Delete(string) error
	GetByUser(uint) ([]models.Segment, error)
}

type User interface {
	// get user's segments
	UserSegments(uint) ([]Segment, error)
}
