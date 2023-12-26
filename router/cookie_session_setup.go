package router

import (
	"crypto/rand"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func setupCookieSession(r *gin.Engine) {
	authKey := make([]byte, 32)
	encriptionKey := make([]byte, 32)
	_, err := rand.Read(authKey)
	_, err = rand.Read(encriptionKey)
	if err != nil {
		panic("Error generating random auth/encryption key")
	}
	store := cookie.NewStore(authKey, encriptionKey)
	store.Options(sessions.Options{MaxAge: 60 * 60 * 24, HttpOnly: true, Secure: false})
	r.Use(sessions.Sessions("userSession", store))
}
