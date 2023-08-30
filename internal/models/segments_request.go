package models

type SegmentRequest struct {
	Add    []Segment `json:"add"`
	Delete []string  `json:"delete"`
}
