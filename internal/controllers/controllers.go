package controllers

import "github.com/gin-gonic/gin"

type Segments interface {
	// create new segment by slug name
	Save(*gin.Context)
}
