package main_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine

	expectedPingResponse    = gin.H{"message": "pong"}
	expectedAPIResponse     = gin.H{"message": "API handler"}
	expectedHealthzResponse = gin.H{"status": "ok"}
	expectedRootResponse    = gin.H{"message": "Welcome to the server!"}
)

func init() {
	// Set up a test environment
	os.Setenv("HOST_ADDR", ":8081")
	gin.SetMode(gin.TestMode)
	router = gin.Default()

	// Register routes
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, expectedPingResponse)
	})
	router.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, expectedAPIResponse)
	})
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, expectedHealthzResponse)
	})
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, expectedRootResponse)
	})
}

func TestMain(t *testing.T) {
	// Test the main function
	t.Run("TestMain", func(t *testing.T) {
		// Define the test cases
		testCases := []struct {
			name           string
			method         string
			path           string
			expectedStatus int
			expectedBody   gin.H
		}{
			{"Ping", "GET", "/ping", http.StatusOK, expectedPingResponse},
			{"API", "GET", "/api", http.StatusOK, expectedAPIResponse},
			{"Healthz", "GET", "/healthz", http.StatusOK, expectedHealthzResponse},
			{"Root", "GET", "/", http.StatusOK, expectedRootResponse},
		}

		// Run the test cases
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				req, err := http.NewRequest(tc.method, tc.path, nil)
				if err != nil {
					t.Fatalf("Failed to create request: %v", err)
				}

				rec := httptest.NewRecorder()
				router.ServeHTTP(rec, req)

				if rec.Code != tc.expectedStatus {
					t.Errorf("Expected status %d, got %d", tc.expectedStatus, rec.Code)
				}

				var body gin.H
				err = json.Unmarshal(rec.Body.Bytes(), &body)
				if err != nil {
					t.Fatalf("Failed to unmarshal response body: %v", err)
				}

				if !reflect.DeepEqual(body, tc.expectedBody) {
					t.Errorf("Expected body %v, got %v", tc.expectedBody, body)
				}
			})
		}
	})
}
