package repositories

import "github.com/burravlev/avito-tech-test/internal/models"

type User interface {
	Save(*models.User) (*models.User, error)
}

type Segment interface {
	Save(*models.Segment) (*models.Segment, error)
}
