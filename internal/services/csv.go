package services

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/burravlev/avito-tech-test/internal/models"
	"github.com/sirupsen/logrus"
)

func marshal(segments []models.Segment) [][]string {
	res := [][]string{}
	for _, s := range segments {
		res = append(res, []string{
			fmt.Sprintf("%d", *s.UserID), s.Name, "added", fmt.Sprintf("%s", s.Added),
		})
		if s.Deleted != nil {
			res = append(res, []string{
				fmt.Sprintf("%d", *s.UserID), s.Name, "deleted", fmt.Sprintf("%s", s.Deleted),
			})
		}
	}
	return res
}

// TODO Put path in config
func createCSV(values [][]string, filename string) error {
	path := "files/" + filename
	f, err := os.Create(path)
	if err != nil {
		logrus.Errorf("services.createCSV: %s", err)
		return err
	}
	w := csv.NewWriter(f)
	defer w.Flush()

	for _, record := range values {
		if err := w.Write(record); err != nil {
			logrus.Errorf("error writing record to file: %s", err)
		}
	}
	return nil
}
