package repositories

import (
	"github.com/burravlev/avito-tech-test/internal/models"
	"gorm.io/gorm"
)

type segment struct {
	db *gorm.DB
}

func SegmentRepository(db *gorm.DB) Segment {
	return &segment{db}
}

func (s *segment) Save(segment *models.Segment) (*models.Segment, error) {
	tx := s.db.Begin()
	result := tx.Model(segment).Save(&segment)
	tx.Commit()
	return segment, result.Error
}
