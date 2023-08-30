package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/burravlev/avito-tech-test/internal/models"
	"github.com/burravlev/avito-tech-test/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type segments struct {
	service services.Segment
}

func SegmentsController(service services.Segment) Segments {
	return &segments{service}
}

//	@Summary		Create segment
//	@Description	create segment
//	@Tags			segments
//	@Accept			json
//	@Produce		json
//	@Param			segment	body		models.Segment	true	"Segment Info"
//	@Success		201		{object}	models.Segment
//	@Failure		400		{object}	models.Error
//	@Failure		500		{object}	models.Error
//	@Router			/api/v1/segments [post]
func (s *segments) Create(c *gin.Context) {
	var body models.Segment
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: "Field name can't be empty"})
		return
	}
	err := s.service.Create(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{Message: "Internal server error"})
	} else {
		c.JSON(http.StatusCreated, body)
	}
}

//	@Summary		Delete segment
//	@Tags			segments
//	@Description	deletes segment (also deletes it from users)
//	@Param			name	path	string	true	"Segment slug name"
//	@Success		204
//	@Failure		500	{object}	models.Error
//	@Router			/api/v1/segments/{name} [delete]
func (s *segments) Delete(c *gin.Context) {
	name := c.Param("name")
	if err := s.service.Delete(name); err != nil {
		log.Errorf("controllers.Delete: %s", err)
		c.JSON(http.StatusInternalServerError, models.Error{Message: "Internal server error"})
		return
	}
	c.Status(http.StatusNoContent)
}

//	@Summary		User's active segments
//	@Tags			users
//	@Description	Gets user's active segments
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	models.User
//	@Failure		404	{object}	models.Error
//	@Failure		500	{object}	models.Error
//	@Router			/api/v1/users/{id}/segments [get]
func (s *segments) GetUserSegments(c *gin.Context) {
	id, ok := extractID(c)
	if !ok {
		return
	}
	user, err := s.service.GetUserSegments(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, models.Error{Message: "User not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{Message: "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, user)
}

//	@Summary		Update user's segments
//	@Tags			users
//	@Description	Adds and deletes user's segments
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int						true	"User ID"
//	@Param			user	body		models.SegmentRequest	true	"Segments update"
//	@Success		200		{object}	models.User
//	@Failure		404		{object}	models.Error
//	@Failure		500		{object}	models.Error
//	@Router			/api/v1/users/{id}/segments [post]
func (s *segments) UpdateSegments(c *gin.Context) {
	id, ok := extractID(c)
	if !ok {
		return
	}
	var body models.SegmentRequest
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	fmt.Println(body)
	user, err := s.service.UpdateSegments(uint(id), body.Add, body.Delete)
	if err != nil {
		log.Errorf("controllers.UpdateSegments: %s", err)
		c.JSON(http.StatusInternalServerError, models.Error{Message: "Internal server error"})
	}
	c.JSON(http.StatusOK, user)
}

//	@Summary		User report .csv file
//	@Tags			users
//	@Description	Deletes user's segments
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int		true	"User ID"
//	@Param			from	query		string	true	"From date"
//	@Param			from	query		string	false	"To date"
//	@Success		200		{object}	models.History
//	@Failure		400		{object}	models.Error
//	@Failure		500		{object}	models.Error
//	@Router			/api/v1/users/{id}/segments/history [get]
func (s *segments) History(c *gin.Context) {
	id, ok := extractID(c)
	if !ok {
		return
	}
	from := c.Query("from")
	to := c.Query("to")
	if from == "" {
		c.JSON(http.StatusBadRequest, models.Error{Message: "Query param \"from\" can't be empty"})
		return
	}
	filename, err := s.service.GenerateReport(uint(id), from, to)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, models.Error{Message: "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"url": "/files/reports/" + filename})
}

func extractID(c *gin.Context) (int, bool) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: "ID must be number"})
		return 0, false
	}
	return id, true
}
