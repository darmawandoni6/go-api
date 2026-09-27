package middleware

import (
	"log"
	"net/http"

	"github.com/donidarmawan/test-go/internal/dto"
	"github.com/gin-gonic/gin"
)

// Recovery catches panics and returns a 500 JSON response instead of crashing.
func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered any) {
		log.Printf("[PANIC] %v", recovered)
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
		})
	})
}

// NoRoute handles requests to undefined routes (404).
func NoRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "route not found",
		})
	}
}

// NoMethod handles requests with an unsupported HTTP method (405).
func NoMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, dto.ErrorResponse{
			Code:    http.StatusMethodNotAllowed,
			Message: "method not allowed",
		})
	}
}
