package controllers

import (
	"net/http"

	"github.com/burravlev/avito-tech-test/internal/models"
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
	var body models.Segment
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	if err := s.service.Save(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, body)
}

func (s *segments) Remove(c *gin.Context) {

}
