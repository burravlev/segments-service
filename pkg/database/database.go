/*
* @File: pkg.database.database.go
* @Desctiption: connection to database
* @Author: Daniil Buravlev (burravlev.d.a@gmail.com)
 */
package database

import (
	"github.com/burravlev/avito-tech-test/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *config.DB) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.URL), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("error connecting to database: %s", err)
	}
	// db.AutoMigrate(&models.SegmentUser{}, &models.Segment{})
	return db, nil
}
