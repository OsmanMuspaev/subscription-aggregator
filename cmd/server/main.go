package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yourusername/subscription-service/internal/db"
)

func main() {
	db.Connect()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	port := "8080"
	logrus.Infof("Server running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		logrus.Fatal(err)
	}
}