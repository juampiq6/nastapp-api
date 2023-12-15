package router

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	store := cookie.NewStore([]byte("secretAuthKey"), []byte("secretEncryptionKey"))
	r.Use(sessions.Sessions("countrySession", store))
	setupSwagger(r)

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/gas_stations", getGasStationsHandler)
	r.POST("/petrol_price/:countryCode", insertPetrolPriceHandler)

	return r
}
