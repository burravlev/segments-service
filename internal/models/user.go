package models

type User struct {
	ID       int64 `db:"id" gorm:"primary_key"`
	Segments Segment
}

func (User) TableName() string {
	return "users"
}
