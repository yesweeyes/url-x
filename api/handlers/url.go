package handlers

import (
	"net/http"
	"os"
	"time"
	"url-x-api/redis"
	"url-x-api/utils"

	"github.com/gin-gonic/gin"
)

type ShortenRequest struct {
	URL string `json:"url" binding:"required"`
}

func Shorten(c *gin.Context) {
	var req ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	code := utils.GenerateCode(6)

	// Store the original URL in Redis with a 24-hour expiration time
	redis.Client.Set(redis.Ctx, code, req.URL, 24*time.Hour)

	// Construct the short URL using the base URL from environment variable
	shortURL := os.Getenv("BASE_URL") + "/" + code

	// Respond with the short URL
	c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
}

func Redirect(c *gin.Context) {
	code := c.Param("code")
	url, err := redis.Client.Get(redis.Ctx, code).Result()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "code not found"})
		return
	}
	c.Redirect(http.StatusMovedPermanently, url)
}
