package repositories

import "github.com/burravlev/avito-tech-test/internal/models"

type Segment interface {
	Create(*models.Segment) error
	Delete(name string) error
	GetByUserID(userId uint) ([]models.Segment, error)
	Update(userId uint, add, delete []string) ([]models.Segment, error)
}
