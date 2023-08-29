package repositories

import "github.com/burravlev/avito-tech-test/internal/models"

type Segment interface {
	// Save creates or updates segment
	Save(*models.Segment) (*models.Segment, error)
}
