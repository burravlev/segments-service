package v1

import "github.com/gin-gonic/gin"

type segmentRoutes struct {
}

func NewSegmentRoutes(r *gin.Engine) {
	r.POST("/api/v1/segments")
}
