package models

type Segment struct {
	ID           uint          `db:"id" gorm:"primary_key" json:"-"`
	Name         string        `db:"name" gorm:"unique" json:"name" binding:"required"`
	SegmentUsers []SegmentUser `gorm:"foreignKey:SegmentID" json:"-"`
}

func (Segment) TableName() string {
	return "segments"
}
