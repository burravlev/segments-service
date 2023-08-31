/*
* @File: pkg.database.database.go
* @Desctiption: connection to database
* @Author: Daniil Buravlev (burravlev.d.a@gmail.com)
 */
package database

import (
	"github.com/burravlev/avito-tech-test/config"
	"github.com/burravlev/avito-tech-test/internal/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(cfg *config.DB) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.URL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		return nil, err
	}
	log.Info("Migrating database...")
	err = db.AutoMigrate(&models.Segment{})
	return db, err
}
