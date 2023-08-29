package repositories

import "github.com/burravlev/avito-tech-test/internal/models"

type Segment interface {
	// Save creates or updates segment
	Save(*models.Segment) error
	// Deletes segment by segment name
	Delete(name string) error
	// Gets active user's segments
	GetByUser(uint) ([]models.Segment, error)
}
