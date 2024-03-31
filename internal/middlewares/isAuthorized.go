package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sandbox-science/deep-focus/internal/utils"
)

// IsAuthorized is a middleware function that checks if the request is authorized.
func IsAuthorized() gin.HandlerFunc {
	var allowedRoutes = []string{"/", "/login", "/signup"}

	return func(c *gin.Context) {
		// Allow access to homepage, /login and /signup routes without token validation
		for _, route := range allowedRoutes {
			if c.Request.URL.Path == route {
				c.Next()
				return
			}
		}

		// Check for token cookie
		cookie, err := c.Cookie("token")
		if err != nil {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		// Validate token
		claims, err := utils.ParseToken(cookie)
		if err != nil {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		// Set role claim in the context
		c.Set("role", claims.Role)
		c.Next()
	}
}
