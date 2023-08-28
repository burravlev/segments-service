package models

type User struct {
	ID int64 `db:"id" gorm:"primary_key"`
}

func (User) TableName() string {
	return "users"
}
