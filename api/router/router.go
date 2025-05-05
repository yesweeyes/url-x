package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"url-x-api/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/api/shorten", handlers.Shorten)

	r.GET("/:code", handlers.Redirect)

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return r
}
