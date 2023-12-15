package router

import (
	countrylocator "nastapp-api/country_locator"
	"nastapp-api/db"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// @Summary      Lists all gas stations near a point restricted to a maximum distance, optionally with the latest prices attached
// @Description  lat and long parameters must be double
// @Tags         gas_stations
// @Produce      json
// @Param        point    query     db.LatLong  true  "point"
// @Param        maxDistance    query     int  true  "Distance in meters"
// @Param		 petrolPrices    query    bool    true    "Add gas prices to each place response"
// @Success      200  {array}   db.PlaceResult
// @Success      200  {array}   db.PlaceResultWPrices
// @Failure      400  {object}  router.APIError
// @Failure      500  {object}  router.APIError
// @Router       /gas_stations [get]
func getGasStationsHandler(c *gin.Context) {
	// The coordinates are parsed
	var coord *db.LatLong
	err := c.ShouldBindQuery(&coord)
	if err != nil {
		e := APIError{Code: 400, Description: "The lat and long arguments were invalid or missing. This arguments must be float numbers."}
		c.AbortWithStatusJSON(e.Code, e)
		return
	}
	// The maxDistance argument is parsed
	maxDist, err := strconv.Atoi(c.Query("maxDistance"))
	if err != nil {
		e := APIError{Code: 400, Description: "The maxDistance argument is invalid or missing"}
		c.AbortWithStatusJSON(e.Code, e)
		return
	}
	// The gas price bool is parsed
	prices, err := strconv.ParseBool(c.Query("petrolPrices"))
	if err != nil {
		e := APIError{Code: 400, Description: "The petrolPrices argument is invalid or missing"}
		c.AbortWithStatusJSON(e.Code, e)
		return
	}
	// Gets country saved in session and passes it to the db handler
	if prices {
		countryCode, err := getCountryFromCookiesOrFetchFromCoord(c, coord)
		if err != nil {
			// The request is already responded with an error, just exit this handler
			return
		}
		results := db.GetGasStationsAndPricesForLocation(coord, &maxDist, countryCode)
		c.JSON(http.StatusOK, *results)
	} else {
		results := db.GetGasStationsForLocation(coord, &maxDist)
		c.JSON(http.StatusOK, *results)
	}
}

func getCountryFromCookiesOrFetchFromCoord(c *gin.Context, coord *db.LatLong) (string, error) {
	// The session is checked for avoiding quering from which country the coordinates correspond to
	session := sessions.Default(c)
	sessionCountryCode := session.Get("countryCode")
	if sessionCountryCode == nil {
		// If no cookies are available, then query the country
		countryCode, err := countrylocator.GetCountryByCoordinates(coord.Lat, coord.Long)
		if err != nil {
			e := APIError{Code: 500, Description: "Error finding country from coordinates"}
			c.AbortWithStatusJSON(e.Code, e)
			return "", e
		}
		session.Set("countryCode", countryCode)
		return countryCode, nil
	} else {
		return sessionCountryCode.(string), nil
	}
}
