package controllers

import "github.com/gin-gonic/gin"

type Segments interface {
	Save(*gin.Context)
}
