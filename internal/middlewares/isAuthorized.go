package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sandbox-science/deep-focus/internal/utils"
)

// IsAuthorized is a middleware function that checks if the request is authorized.
// It verifies the presence and validity of a token cookie, and sets the role claim in the context.
// If the token is missing or invalid, it returns a 401 Unauthorized response.
func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")

		if err != nil {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(cookie)

		if err != nil {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		c.Set("role", claims.Role)
		c.Next()
	}
}
