/*
* @File: pkg.database.database.go
* @Desctiption: connection to database
* @Author: Daniil Buravlev (burravlev.d.a@gmail.com)
 */
package database

import (
	"fmt"

	"github.com/burravlev/avito-tech-test/config"
	"github.com/burravlev/avito-tech-test/internal/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *config.DB) (*gorm.DB, error) {
	fmt.Println(cfg.URL)
	db, err := gorm.Open(postgres.Open(cfg.URL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Info("Migrating database...")
	err = db.AutoMigrate(&models.Segment{}, &models.User{})
	return db, err
}
