package middleware

import (
	"net/http"
	"strings"

	"github.com/donidarmawan/test-go/internal/config"
	"github.com/donidarmawan/test-go/internal/utils"
	"github.com/gin-gonic/gin"
)

const UserIDKey = "user_id"
const UserEmailKey = "user_email"

// AuthGuard validates the Bearer JWT token from the Authorization header.
func AuthGuard(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization format, expected: Bearer <token>"})
			return
		}

		claims, err := utils.ValidateToken(parts[1], cfg.JWTSecret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		// Inject claims into context for downstream handlers
		c.Set(UserIDKey, claims.UserID.String())
		c.Set(UserEmailKey, claims.Email)

		c.Next()
	}
}
