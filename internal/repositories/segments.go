package repositories

import (
	"github.com/burravlev/avito-tech-test/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type segment struct {
	db *gorm.DB
}

func SegmentRespository(db *gorm.DB) Segment {
	return &segment{db}
}

func (s *segment) Create(segment *models.Segment) error {
	err := s.db.Create(&segment).Error
	return err
}

func (s *segment) Delete(name string) error {
	return s.db.Exec("DELETE FROM segments WHERE name = ?", name).Error
}

func (s *segment) GetByUserID(userId uint) ([]models.Segment, error) {
	var segments []models.Segment
	err := s.db.Model(&models.Segment{}).Where("user_id = ?", userId).Find(&segments).Error
	return segments, err
}

func (s *segment) Update(userId uint, add, delete []string) ([]models.Segment, error) {
	tx := s.db.Begin()
	// setting transation isolation level
	err := tx.Exec("set transaction isolation level repeatable read").Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	segments := segmentsFromNames(userId, add)
	err = tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&segments).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Exec("DELETE FROM segments WHERE user_id = ? and name in ?", userId, delete).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Where("user_id = ?", userId).Find(&segments).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return segments, nil
}

func segmentsFromNames(userId uint, names []string) []models.Segment {
	segments := make([]models.Segment, len(names))
	for i := range segments {
		segments[i] = models.Segment{UserID: &userId, Name: names[i]}
	}
	return segments
}
