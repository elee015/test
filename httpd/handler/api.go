package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorld() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, gin.H{
			"data": "Hello World",
		})
	}
}
