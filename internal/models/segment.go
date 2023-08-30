package models

import "time"

type Segment struct {
	ID      uint       `db:"id" gorm:"primary_key" json:"-" csv:"-"`
	UserID  *uint      `db:"user_id" gorm:"uniqueIndex:idx_user_segment" json:"-" csv:"user_id"`
	Name    string     `db:"name" gorm:"uniqueIndex:idx_user_segment" json:"name" binding:"required" csv:"name"`
	Added   *time.Time `db:"added" gorm:"default:now()" json:"-" csv:"added"`
	Deleted *time.Time `db:"deleted" json:"delete_at,omitempty" csv:"deleted"`
	PerCent *int       `db:"per_cent" json:"per_cent,omitempty"`
}

type AddedUser struct {
	ID uint `db:"id" gorm:"primary_key"`
}

func (Segment) TableName() string {
	return "segments"
}
