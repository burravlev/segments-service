package models

type Segment struct {
	ID   uint   `db:"id" gorm:"primary_key" json:"-"`
	Name string `db:"name" gorm:"unique" json:"name" binding:"required"`
}

func (Segment) TableName() string {
	return "segments"
}
