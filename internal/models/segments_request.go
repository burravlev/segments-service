package models

type SegmentRequest struct {
	UserID           uint     `json:"user_id" binding:"required"`
	SegmentsToAdd    []string `json:"segments_to_add; ommitempty"`
	SegmentsToDelete []string `json:"segments_to_delete; ommitempty"`
}
