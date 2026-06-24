package main

import (
	"net/http"

	"github.com/chukwuka4u/clickjson-server/services"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/api/mock", services.HandleMockGeneration)

	router.Run("0.0.0.0:8080")
}
