package main

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/yourusername/subscription-service/swagger"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yourusername/subscription-service/internal/config"
	"github.com/yourusername/subscription-service/internal/db"
	"github.com/yourusername/subscription-service/internal/handler"
	"github.com/yourusername/subscription-service/internal/repository"
)

func main() {
	cfg := config.LoadConfig()
	db.Connect(cfg)

	r := gin.Default()

	subRepo := repository.NewSubscriptionRepo()
	subHandler := handler.NewSubscriptionHandler(subRepo)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/subscriptions", subHandler.Create)
	r.GET("/subscriptions", subHandler.List)
	r.GET("/subscriptions/summary", subHandler.Summary)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := "8080"
	logrus.Infof("Server running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		logrus.Fatal(err)
	}
}