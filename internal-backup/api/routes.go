package api

import (
	"github.com/gin-gonic/gin"
	"local/web-service-gin/internal/service"
)

func RegisterRoutes(r *gin.Engine, svc *service.DownloadService) {
	r.GET("/download", func(c *gin.Context) {
		svc.StreamExcel(c.Writer, c.Request)
	})
}
