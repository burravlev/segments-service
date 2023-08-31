package repositories

import (
	"time"

	"github.com/burravlev/avito-tech-test/internal/models"
)

type Segment interface {
	// creates segment in database
	Create(*models.Segment) error
	// deletes segments from database
	Delete(name string) error
	// gets user's segments
	GetByUserID(userId uint) ([]models.Segment, error)
	// updates user's segments
	Update(userId uint, add []models.Segment, delete []string) ([]models.Segment, error)
	// gets user's segment in time interval
	GetInInterval(userId uint, from, to time.Time) ([]models.Segment, error)
	// creates segment with rendom percentage of all users in database
	CreateWithAutoSplit(segment *models.Segment) error
}
