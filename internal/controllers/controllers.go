package controllers

import "github.com/gin-gonic/gin"

type Segements interface {
	Create(*gin.Context)
}
