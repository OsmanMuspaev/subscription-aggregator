package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yourusername/subscription-service/internal/db"
	"github.com/yourusername/subscription-service/internal/handler"
	"github.com/yourusername/subscription-service/internal/repository"
)

func main() {
	db.Connect()

	r := gin.Default()

	subRepo := repository.NewSubscriptionRepo()
	subHandler := handler.NewSubscriptionHandler(subRepo)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/subscriptions", subHandler.Create)
	r.GET("/subscriptions", subHandler.List)
	r.GET("/subscriptions/summary", subHandler.Summary)

	port := "8080"
	logrus.Infof("Server running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		logrus.Fatal(err)
	}
}