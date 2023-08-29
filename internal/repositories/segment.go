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

func (s *segment) Save(segment *models.Segment) error {
	return s.db.Model(segment).Save(&segment).Error
}

func (s *segment) Delete(name string) error {
	tx := s.db.Begin()
	err := tx.Exec("set transaction isolation level repeatable read").Error
	if err != nil {
		tx.Rollback()
		return err
	}
	var segment models.Segment
	err = tx.Model(&models.Segment{}).Where("name = ?", name).Scan(&segment).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Select("SegmentUsers").Delete(&segment).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}
