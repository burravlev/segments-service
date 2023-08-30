package services

import (
	"fmt"
	"os"
	"time"

	"github.com/burravlev/avito-tech-test/internal/models"
	"github.com/burravlev/avito-tech-test/internal/repositories"
	"github.com/gocarina/gocsv"
)

type segment struct {
	repository repositories.Segment
}

func SegmentService(repository repositories.Segment) Segment {
	return &segment{repository}
}

func (s *segment) Create(segment *models.Segment) error {
	return s.repository.Create(segment)
}

func (s *segment) Delete(name string) error {
	return s.repository.Delete(name)
}

func (s *segment) GetUserSegments(userId uint) (*models.User, error) {
	segments, err := s.repository.GetByUserID(userId)
	return &models.User{ID: userId, Segments: segments}, err
}

func (s *segment) UpdateSegments(userId uint, add []models.Segment, delete []string) (*models.User, error) {
	segments, err := s.repository.Update(userId, add, delete)
	return &models.User{ID: userId, Segments: segments}, err
}

func (s *segment) GenerateReport(userId uint, from, to string) (string, error) {
	f, t, err := parseTime(from, to)
	if err != nil {
		return "", err
	}
	segments, err := s.repository.GetInInterval(userId, f, t)
	filename := fmt.Sprintf(`%d-%s-%s.csv`, userId, from, to)
	err = createCSV(segments, filename)
	return filename, err
}

func parseTime(from, to string) (time.Time, time.Time, error) {
	layout := "2006-01"
	var f, t time.Time
	var err error
	f, err = time.Parse(layout, from)
	if err != nil {
		return f, t, err
	}
	if to == "" {
		t = time.Now()
	} else {
		t, err = time.Parse(layout, to)
		if err != nil {
			return t, f, err
		}
	}
	return f, t, nil
}

func createCSV(segments []models.Segment, filename string) error {
	path := "files\\" + filename
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	return gocsv.MarshalFile(&segments, file)
}
