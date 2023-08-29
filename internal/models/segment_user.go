package models

type SegmentUser struct {
	ID        uint `db:"id" gorm:"primarykey"`
	UserID    uint `db:"user_id"`
	SegmentID uint
}
