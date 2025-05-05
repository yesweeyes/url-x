package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"url-x-api/redis"
	"url-x-api/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found. Using system environment variables.")
	}

	redis.InitRedis()

	r := router.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server starting on port", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
