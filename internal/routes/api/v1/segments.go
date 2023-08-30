package v1

import (
	"github.com/burravlev/avito-tech-test/internal/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SegmentRoutes(r *gin.Engine, c controllers.Segments) {
	r.POST("/api/v1/segments", c.Create)
	r.DELETE("/api/v1/segments/:name", c.Delete)
	r.POST("/api/v1/segments/:name")

	r.GET("/api/v1/users/:id/segments", c.GetUserSegments)
	r.POST("/api/v1/users/:id/segments", c.UpdateSegments)
	r.GET("/api/v1/users/:id/segments/history", c.History)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
