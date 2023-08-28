package repositories

import "github.com/burravlev/avito-tech-test/internal/models"

type user struct {
}

func (u *user) Save(*models.User) (*models.User, error) {
	return nil, nil
}
