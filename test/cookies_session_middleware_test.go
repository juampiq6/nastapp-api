package test

import (
	"encoding/json"
	"nastapp-api/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCookieSessionMiddleware(t *testing.T) {
	const userIdValue = "123"
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.Use(router.CookieSessionMiddleware())

	r.GET("/set", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("userID", userIdValue)
		session.Save()
		c.JSON(http.StatusOK, gin.H{"message": "User ID set in session"})
	})

	r.GET("/get", func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("userID")

		c.JSON(http.StatusOK, gin.H{"userID": userID})
	})

	// Create a test request to set the user ID
	setReq, err := http.NewRequest(http.MethodGet, "/set", nil)
	assert.NoError(t, err)

	setRecorder := httptest.NewRecorder()
	r.ServeHTTP(setRecorder, setReq)

	// Check the cookies in the response exist
	sessionCookie := setRecorder.Result().Cookies()[0]
	assert.NotNil(t, *sessionCookie)
	assert.Equal(t, http.StatusOK, setRecorder.Code)

	// Create a test request to get the user ID
	getReq, err := http.NewRequest(http.MethodGet, "/get", nil)
	assert.NoError(t, err)
	// Set the cookie header using the set request response cookie
	clientCookieHeader := sessionCookie.Name + "=" + sessionCookie.Value
	getReq.Header.Set("cookie", clientCookieHeader)
	// Make the get request
	getRecorder := httptest.NewRecorder()
	r.ServeHTTP(getRecorder, getReq)
	assert.Equal(t, http.StatusOK, getRecorder.Code)

	var getResponse map[string]interface{}

	assert.NoError(t, json.Unmarshal(getRecorder.Body.Bytes(), &getResponse))
	// Check the user ID retrieved from the session
	assert.Equal(t, userIdValue, getResponse["userID"])
}
