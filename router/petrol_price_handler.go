package router

import (
	"encoding/json"
	countrylocator "nastapp-api/country_locator"
	"nastapp-api/db"
	"nastapp-api/db/petrolprices"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary      Upload specific country petrol prices for a station id
// @Description  The param
// @Tags         gas_price
// @Produce      plain
// @Param		 countryCode    path    string    true    "Country code"
// @Param        gasPrice    body     gasprices.GasPriceArgentina  true  "The param type GasPriceArgentina is not mandatory, depends on the country pass in the query params. That is the case for Argentina"
// @Success      200  {string}
// @Failure      400  {object}  router.APIError
// @Failure      500  {object}  router.APIError
// @Router       /gas_prices [post]
func insertPetrolPriceHandler(c *gin.Context) {
	countryCode := strings.ToUpper(c.Params.ByName("countryCode"))
	var parsingError error
	var price interface{}
	switch countryCode {
	case countrylocator.Argentina:
		var p petrolprices.PetrolPriceArgentina
		c.BindJSON(&p)
		p.Timestamp = time.Now()
		price = &p
		_, parsingError = json.Marshal(p)
		if parsingError != nil {
			break
		}
	default:
		e := APIError{Code: 400, Description: "Error: the parsing method for this station price's country is not implemented"}
		c.AbortWithStatusJSON(e.Code, e)
		return
	}
	if parsingError != nil {
		er := APIError{Code: 400, Description: "Error parsing gas price model.\n" + parsingError.Error()}
		c.AbortWithStatusJSON(er.Code, er.Description)
		return
	}
	id, e := db.InsertGasPriceForStation(&price, countryCode)
	if e != nil {
		er := APIError{Code: 500, Description: "Error uploading record to DB.\n" + e.Error()}
		c.AbortWithStatusJSON(er.Code, er.Description)
		return
	}
	c.String(201, id)
}
