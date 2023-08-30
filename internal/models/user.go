package models

type User struct {
	ID       uint      `json:"id"`
	Segments []Segment `json:"segments"`
}
