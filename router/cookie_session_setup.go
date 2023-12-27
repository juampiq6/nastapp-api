package router

import (
	"crypto/rand"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func cookieSessionMiddleware() gin.HandlerFunc {
	authKey := make([]byte, 32)
	encriptionKey := make([]byte, 32)
	_, err := rand.Read(authKey)
	_, err = rand.Read(encriptionKey)
	if err != nil {
		panic("Error generating random auth/encryption key")
	}
	store := cookie.NewStore(authKey, encriptionKey)
	store.Options(sessions.Options{MaxAge: 60 * 60 * 24, HttpOnly: true, Secure: false})
	return sessions.Sessions("userSession", store)
}
