package controllers

import "github.com/gin-gonic/gin"

type Segments interface {
	// create new segment by slug name
	Save(*gin.Context)
	// deletes segment by slug name
	Delete(*gin.Context)
	// get all user's active segments
	UserSegments(*gin.Context)
}
