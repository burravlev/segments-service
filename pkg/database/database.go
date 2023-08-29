/*
* @File: pkg.database.database.go
* @Desctiption: connection to database
* @Author: Daniil Buravlev (burravlev.d.a@gmail.com)
 */
package database

import (
	"github.com/burravlev/avito-tech-test/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *config.DB) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.URL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// db.AutoMigrate(&models.SegmentUser{}, &models.Segment{})
	return db, nil
}
