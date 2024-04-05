package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sandbox-science/deep-focus/internal/utils"
)

/*
JwtAuthMiddleware checks if the user is
authorized to access the protected routes.
*/
func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := utils.ValidateToken(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"Unauthorized": "Authentication required"})
			fmt.Println(err)
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
