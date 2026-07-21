package middleware

import (
	"net/http"
	"strings"

	"github.com/Bhandavya345/Employee-Management/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates JWT token
func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		// Get Authorization header
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Authorization header missing",
			})
			c.Abort()
			return
		}

		// Check Bearer Token format
		tokenParts := strings.Split(authHeader, " ")

		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Invalid Authorization format",
			})
			c.Abort()
			return
		}

		token := tokenParts[1]

		// Validate JWT
		claims, err := utils.ValidateJWT(token)

		if err != nil {

			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		// Store user information in Gin context
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)

		c.Next()
	}
}
