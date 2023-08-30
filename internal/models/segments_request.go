package models

type SegmentRequest struct {
	Add    []string `json:"add"`
	Delete []string `json:"delete"`
}
