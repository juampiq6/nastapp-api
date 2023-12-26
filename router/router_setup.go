package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const gas_station_path = "/gas_stations"
const petrol_price_path = "/petrol_price"

func SetupRouter() *gin.Engine {

	r := gin.Default()
	setupCookieSession(r)
	setupSwagger(r)

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET(gas_station_path, getGasStationsHandler)
	r.POST(petrol_price_path+"/:countryCode", insertPetrolPriceHandler)

	return r
}
