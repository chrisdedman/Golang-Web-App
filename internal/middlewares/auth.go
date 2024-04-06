package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-web-app/internal/utils"
)

/*
AuthMiddleware checks if the user is
authorized to access the protected routes.
*/
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := utils.ValidateToken(c)

		if err != nil {
			c.HTML(http.StatusUnauthorized, "errors.html", gin.H{
				"Unauthorized": "Authentication required",
			})
			fmt.Println(err)
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
