package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/burravlev/avito-tech-test/internal/models"
	"github.com/burravlev/avito-tech-test/internal/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
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

func (s *segments) Delete(c *gin.Context) {
	name := c.Param("name")
	err := s.service.Delete(name)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, models.Error{Message: "server error occured"})
		return
	}
	c.Status(http.StatusNoContent)
}

func (s *segments) UserSegments(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	fmt.Println(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	segments, err := s.service.GetByUser(uint(id))
	if errors.Is(gorm.ErrRecordNotFound, err) {
		c.JSON(http.StatusNotFound, models.Error{Message: "User doesn't exist"})
		return
	}
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, models.Error{Message: "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, segments)
}
