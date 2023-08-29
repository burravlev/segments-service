package v1

import (
	"github.com/burravlev/avito-tech-test/internal/controllers"
	"github.com/gin-gonic/gin"
)

type segmentRoutes struct {
}

func SegmentRoutes(r *gin.Engine, c controllers.Segments) {
	r.POST("/api/v1/segments", c.Save)
}
