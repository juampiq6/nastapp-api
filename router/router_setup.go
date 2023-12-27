package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const gas_station_path = "/gas_stations"
const petrol_price_path = "/petrol_price"

func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.ForwardedByClientIP = true
	r.Use(RateLimitingMiddleware(time.Second, 6))
	r.Use(RateLimitingMiddleware(time.Minute, 100))
	r.Use(RateLimitingMiddleware(time.Hour*24, 1500))
	r.Use(cookieSessionMiddleware())
	setupSwagger(r)

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET(gas_station_path, getGasStationsHandler)
	r.POST(petrol_price_path+"/:countryCode", insertPetrolPriceHandler)

	return r
}
