package controllers

import (
	"github.com/burravlev/avito-tech-test/internal/services"
	"github.com/gin-gonic/gin"
)

type segments struct {
	service services.Segment
}

func SegmentController(s services.Segment) Segments {
	return &segments{s}
}

func (s *segments) Save(c *gin.Context) {

}
