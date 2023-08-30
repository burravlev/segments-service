package repositories

import (
	"time"

	"github.com/burravlev/avito-tech-test/internal/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// default transaction retries
var retries = 3

type segment struct {
	db *gorm.DB
}

func SegmentRespository(db *gorm.DB) Segment {
	return &segment{db}
}

// saves segment if not exists
func (s *segment) Create(segment *models.Segment) error {
	tx := s.db.Begin()
	err := tx.Exec("set transaction isolation level repeatable read").Error
	if err != nil {
		tx.Rollback()
		return err
	}
	var exists bool
	err = tx.Model(&models.Segment{}).Select("count(*) > 0").Where("name = ?", segment.Name).Find(&exists).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if exists {
		return nil
	}
	err = s.db.Create(&segment).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s *segment) Delete(name string) error {
	return s.db.Exec("UPDATE segments SET deleted = now() WHERE name = ?", name).Error
}

func (s *segment) GetByUserID(userId uint) ([]models.Segment, error) {
	var segments []models.Segment
	err := s.db.Model(&models.Segment{}).Where("user_id = ? and (current_timestamp < deleted or deleted is null)", userId).Find(&segments).Error
	return segments, err
}

func (s *segment) Update(userId uint, add []models.Segment, delete []string) ([]models.Segment, error) {
	tx := s.db.Begin()
	// setting transation isolation level
	err := tx.Exec("set transaction isolation level repeatable read").Error
	if err != nil {
		tx.Rollback()
		logrus.Error(err)
		return nil, err
	}
	for i := range add {
		add[i].UserID = &userId
	}
	err = tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&add).Error
	if err != nil {
		tx.Rollback()
		logrus.Error(err)
		return nil, err
	}
	err = tx.
		Exec("UPDATE segments SET deleted = now() WHERE user_id = ? AND name in ?", userId, delete).
		Error
	if err != nil {
		tx.Rollback()
		logrus.Error(err)
		return nil, err
	}
	err = tx.
		Where("user_id = ? and (current_timestamp < deleted or deleted is null)", userId).
		Find(&add).Error
	if err != nil {
		tx.Rollback()
		logrus.Error(err)
		return nil, err
	}
	err = tx.Commit().Error
	logrus.Error(err)
	return add, err
}

func (s *segment) GetInInterval(userId uint, from, to time.Time) ([]models.Segment, error) {
	var segments []models.Segment
	err := s.db.Model(&models.Segment{}).
		Where("user_id = ? and added >= ? and added <= ?", userId, from, to).Order("added, deleted").
		Find(&segments).Error
	return segments, err
}
