package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	authToken = "sampleToken"
)

type RequestData struct {
	SomeField string `json:"someField" binding:"required"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// Middleware для проверки авторизации
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != authToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{"Unauthorized"})
			return
		}
		c.Next()
	}
}

// Обработчик запросов для /endpoint
func endpointHandler(c *gin.Context) {
	var requestData RequestData
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "received", "data": requestData})
}

func main() {
	router := gin.Default()

	apiGroup := router.Group("/").Use(AuthMiddleware())
	apiGroup.POST("endpoint", endpointHandler)

	router.Run(":3000")
}
