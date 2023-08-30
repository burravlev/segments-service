package models

type User struct {
	ID       uint      `json:"id" swaggertype:"integer" example:"13214"`
	Segments []Segment `json:"segments,omitempty" swaggertype:"object"`
}
