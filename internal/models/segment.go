package models

import "time"

type Segment struct {
	ID      uint       `db:"id" gorm:"primary_key" json:"-" swaggerignore:"true"`
	UserID  *uint      `db:"user_id" gorm:"uniqueIndex:idx_user_segment" json:"-" swaggerignore:"true"`
	Name    string     `db:"name" gorm:"uniqueIndex:idx_user_segment" json:"name" binding:"required" swaggertype:"string" example:"AVITO_TEST_VOICES"`
	Added   *time.Time `db:"added" gorm:"default:now()" json:"-" swaggerignore:"true"`
	Deleted *time.Time `db:"deleted" json:"delete_at,omitempty" swaggertype:"string" example:"2012-10-02T10:00:00-05:00"`
	PerCent *int       `db:"per_cent" json:"per_cent,omitempty" swaggertype:"integer" example:"33"`
}

type AddedUser struct {
	ID uint `db:"id" gorm:"primary_key"`
}

func (Segment) TableName() string {
	return "segments"
}
