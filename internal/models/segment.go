package models

type Segment struct {
	ID     uint   `db:"id" gorm:"primary_key" json:"-"`
	UserID *uint  `db:"user_id" gorm:"uniqueIndex:idx_user_segment" json:"-"`
	Name   string `db:"name" gorm:"uniqueIndex:idx_user_segment" json:"name" binding:"required"`
}

func (Segment) TableName() string {
	return "segments"
}
