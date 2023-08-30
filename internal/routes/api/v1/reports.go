package v1

import (
	"github.com/gin-gonic/gin"
)

//	@Summary		Download .csv file
//	@Tags			files
//	@Description	Downloads .csv file
//	@Produce		text/csv
//	@Param			name	path	string	true	"filename"
//	@Success		200
//	@Failure		404	{object}	models.Error
//	@Failure		500	{object}	models.Error
//	@Router			/files/reports/{name} [post]
func FileRoutes(r *gin.Engine) {
	r.GET("/files/reports/:name", func(c *gin.Context) {
		param := c.Param("name")
		path := "files\\" + param
		c.Header("Content-Description", "attachment; filename="+param)
		c.Header("Content-Type", "text/csv")
		c.File(path)
	})
}
