package v1

import (
	"github.com/gin-gonic/gin"
)

func FileRoutes(r *gin.Engine) {
	r.GET("/files/reports/:name", func(c *gin.Context) {
		param := c.Param("name")
		path := "files\\" + param
		c.Header("Content-Description", "attachment; filename="+param)
		c.Header("Content-Type", "text/csv")
		c.File(path)
	})
}
