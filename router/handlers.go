package router

import (
	"nastapp-api/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary      Lists all gas stations near a point restricted to a maximum distance
// @Description  lat and long parameters must be double
// @Tags         gas_stations
// @Produce      json
// @Param        point    query     db.LatLong  true  "point"
// @Param        maxDistance    query     int  true  "Distance in meters"
// @Success      200  {array}   db.PlaceResult
// @Failure      400  {object}  router.APIError
// @Failure      500  {object}  router.APIError
// @Router       /gas_stations [get]
func getGasStationsHandler(c *gin.Context) {
	var cord *db.LatLong
	err := c.ShouldBindQuery(&cord)
	if err != nil {
		e := APIError{Code: 400, Description: "The lat and long arguments were invalid or missing. This arguments must be float numbers."}
		c.AbortWithStatusJSON(e.Code, e)
		return
	}
	maxDist, err := strconv.Atoi(c.Query("maxDistance"))
	if err != nil {
		e := APIError{Code: 400, Description: "The maxDistance argument is invalid or missing"}
		c.AbortWithStatusJSON(e.Code, e)
		return
	}
	results := db.GetGasStationsForLocation(cord, &maxDist)
	c.JSON(http.StatusOK, *results)
}
