package services

import "github.com/burravlev/avito-tech-test/internal/models"

type Segment interface {
	// creates new segment
	Create(*models.Segment) error
	// deletes segment
	Delete(name string) error
	// gets user's segments
	GetUserSegments(uint) (*models.User, error)
	// updates user's segments
	UpdateSegments(userId uint, add []models.Segment, delete []string) (*models.User, error)
	// generate report
	GenerateReport(userId uint, from, to string) (string, error)
}
