package models

type Segment struct {
	ID           uint          `db:"id" gorm:"primarykey" json:"-"`
	Name         string        `db:"name" gorm:"unique" json:"name"`
	SegmentUsers []SegmentUser `gorm:"foreignKey:SegmentID" json:"-"`
}

func (Segment) TableName() string {
	return "segemnts"
}
