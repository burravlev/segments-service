package models

type Segment struct {
	ID   int64  `db:"id" gorm:"primary_key"`
	Name string `db:"name" gorm:"unique"`
}

func (Segment) TableName() string {
	return "segemnts"
}
