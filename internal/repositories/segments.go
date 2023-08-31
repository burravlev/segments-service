package repositories

import (
	"math/rand"
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
	err = tx.Model(&models.Segment{}).Select("count(*) > 0").
		Where("name = ?", segment.Name).Find(&exists).Error
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

func (s *segment) CreateWithAutoSplit(segment *models.Segment) error {
	var users []uint
	tx := s.db.Begin()
	err := tx.Exec("set transaction isolation level repeatable read").Error
	if err != nil {
		return err
	}
	err = tx.Model(&models.Segment{}).Distinct("user_id").Where("user_id is not null").Find(&users).Error
	segments := AutoSegments(users, *segment.PerCent, segment.Name)
	err = s.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&segments).Error
	return nil
}

// generate list of segments with users
func AutoSegments(users []uint, perCent int, name string) []models.Segment {
	Shuffle(users)
	n := len(users) * perCent / 100
	segments := make([]models.Segment, n)
	for i := 0; i < n; i++ {
		segments[i] = models.Segment{UserID: &users[i], Name: name}
	}
	return segments
}

func Shuffle(arr []uint) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func (s *segment) Delete(name string) error {
	return s.db.Exec("DELETE FROM segments WHERE name = ?", name).Error
}

func (s *segment) GetByUserID(userId uint) ([]models.Segment, error) {
	var segments []models.Segment
	err := s.db.Model(&models.Segment{}).Where("user_id = ? and name is not null and name != ? and (current_timestamp < deleted or deleted is null)", userId, "").Find(&segments).Error
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
		Where("user_id = ? and name is not null and name != ? and (current_timestamp < deleted or deleted is null)", userId, "").
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
