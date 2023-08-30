package v1

import (
	"github.com/burravlev/avito-tech-test/internal/controllers"
	"github.com/gin-gonic/gin"
)

func SegmentRoutes(r *gin.Engine, c controllers.Segments) {
	r.POST("/api/v1/segments", c.Create)
	r.DELETE("/api/v1/segments/:name", c.Delete)
	r.GET("/api/v1/users/:id/segments", c.GetUserSegments)
	r.POST("/api/v1/users/:id/segments", c.UpdateSegments)
}
