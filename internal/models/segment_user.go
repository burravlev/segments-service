package models

type SegmentUser struct {
	ID        uint `db:"id" gorm:"primary_key"`
	UserID    uint `db:"user_id" binding:"required"`
	SegmentID uint
}

func (SegmentUser) TableName() string {
	return "segment_users"
}
