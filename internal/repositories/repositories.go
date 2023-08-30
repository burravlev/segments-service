package repositories

import (
	"time"

	"github.com/burravlev/avito-tech-test/internal/models"
)

type Segment interface {
	Create(*models.Segment) error
	Delete(name string) error
	GetByUserID(userId uint) ([]models.Segment, error)
	Update(userId uint, add []models.Segment, delete []string) ([]models.Segment, error)
	GetInInterval(userId uint, from, to time.Time) ([]models.Segment, error)
}
