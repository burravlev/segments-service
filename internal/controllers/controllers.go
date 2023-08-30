package controllers

import "github.com/gin-gonic/gin"

type Segments interface {
	// create segment
	Create(*gin.Context)
	// deletes segment
	Delete(*gin.Context)
	// gets active user's segments
	GetUserSegments(*gin.Context)
	// update user's active segments
	UpdateSegments(*gin.Context)
}
