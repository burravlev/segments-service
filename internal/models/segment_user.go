package models

type SegmentUser struct {
	ID        uint `db:"id" gorm:"primarykey"`
	UserID    uint `db:"user_id"`
	SegmentID uint
}

func (SegmentUser) TableName() string {
	return "segment_users"
}
