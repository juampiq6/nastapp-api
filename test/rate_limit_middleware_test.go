package test

import (
	"nastapp-api/router"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var duration time.Duration
var limit int

func TestRateLimitingMiddleware(t *testing.T) {
	duration = time.Second
	limit = 6
	t.Run("RateLimitForSecond", pingEndpointWithRateLimit)
	duration = time.Minute
	limit = 100
	t.Run("RateLimitForMinute", pingEndpointWithRateLimit)
	duration = time.Hour * 24
	limit = 1500
	t.Run("RateLimitForDay", pingEndpointWithRateLimit)
}

func pingEndpointWithRateLimit(t *testing.T) {
	start := time.Now()
	router := setupRouterWithRateLimits(duration, limit)
	// Run multiple requests within the rate limit
	for i := 0; i <= limit; i++ {

		req, err := http.NewRequest(http.MethodGet, "/ping", nil)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		elapsedTime := time.Now().Sub(start)

		// Check the status code and response body
		if i < limit {
			assert.Less(t, i, limit)
			assert.Equal(t, http.StatusOK, w.Code)
			assert.LessOrEqual(t, elapsedTime, duration)
		} else {
			assert.GreaterOrEqual(t, i, limit)
			assert.Equal(t, http.StatusTooManyRequests, w.Code)
			assert.LessOrEqual(t, elapsedTime, duration)
		}
	}
}

func setupRouterWithRateLimits(duration time.Duration, limit int) *gin.Engine {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.ForwardedByClientIP = true

	// Apply rate limiting middleware
	r.Use(router.RateLimitingMiddleware(duration, limit))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	return r
}
