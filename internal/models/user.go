package models

type User struct {
	ID       uint      `db:"id" gorm:"primary_key"`
	UserID   uint      `db:"user_id" gorm:"unique;not null"`
	Segments []Segment `gorm:"many2many:user_segments"`
}

func (User) TableName() string {
	return "users"
}
