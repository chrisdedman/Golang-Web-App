package routes_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"testing"

	"github.com/golang-web-app/internal/routes"
	"github.com/stretchr/testify/assert"
)

func TestAuthRoutes(t *testing.T) {
	router := gin.Default()
	db := &gorm.DB{}

	routes.AuthRoutes(router, db)
	assert.NotNil(t, router)

	routes := router.Routes()
	assert.Len(t, routes, 14)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/healthcheck", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"status":"ok"}`, w.Body.String())

	// Test the routes
	assert.Equal(t, "/", routes[0].Path)
	assert.Equal(t, "GET", routes[0].Method)

	assert.Equal(t, "/user/dashboard", routes[1].Path)
	assert.Equal(t, "GET", routes[1].Method)

	assert.Equal(t, "/user/delete", routes[2].Path)
	assert.Equal(t, "GET", routes[2].Method)

	assert.Equal(t, "/user/api", routes[3].Path)
	assert.Equal(t, "GET", routes[3].Method)

	assert.Equal(t, "/user/update", routes[4].Path)
	assert.Equal(t, "GET", routes[4].Method)

	assert.Equal(t, "/user/logout", routes[5].Path)
	assert.Equal(t, "GET", routes[5].Method)

	assert.Equal(t, "/login", routes[6].Path)
	assert.Equal(t, "GET", routes[6].Method)

	assert.Equal(t, "/signup", routes[7].Path)
	assert.Equal(t, "GET", routes[7].Method)

	assert.Equal(t, "/healthcheck", routes[8].Path)
	assert.Equal(t, "GET", routes[8].Method)

	assert.Equal(t, "/login", routes[9].Path)
	assert.Equal(t, "POST", routes[9].Method)

	assert.Equal(t, "/signup", routes[10].Path)
	assert.Equal(t, "POST", routes[10].Method)

	assert.Equal(t, "/user/logout", routes[11].Path)
	assert.Equal(t, "POST", routes[11].Method)

	assert.Equal(t, "/user/update/:user_id", routes[12].Path)
	assert.Equal(t, "PUT", routes[12].Method)

	assert.Equal(t, "/user/delete/:user_id", routes[13].Path)
	assert.Equal(t, "DELETE", routes[13].Method)
}
