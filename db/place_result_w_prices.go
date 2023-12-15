package db

import "nastapp-api/db/petrolprices"

// P type should be PetrolPriceCountry
type PlaceResultWPrices[P any] struct {
	MongoId string `bson:"_id" json:"_id"`
	// GoogleMapsId, not useful for frontend
	// Id                string            `bson:"id" json:"id"`
	DisplayName       string            `bson:"displayName" json:"displayName"`
	AddressComponents AddressComponents `bson:"addressComponents" json:"addressComponents"`
	// FormattedAddress  string            `bson:"formattedAddress" json:"formattedAddress"`
	GoogleMapsUri string  `bson:"googleMapsUri" json:"googleMapsUri"`
	Location      LatLong `bson:"location" json:"location"`
	// LocationGEOJson       LatLongGEOJson    `bson:"locationGEOJson" json:"locationGEOJson"`
	ShortFormattedAddress string   `bson:"shortFormattedAddress" json:"shortFormattedAddress"`
	WeekdayDescriptions   []string `bson:"weekdayDescriptions" json:"weekdayDescriptions"`
	PhoneNumber           string   `bson:"nationalPhoneNumber" json:"nationalPhoneNumber"`
	Rating                string   `bson:"rating" json:"rating"`
	UserRatingCount       string   `bson:"userRatingCount" json:"userRatingCount"`
	// This array is always of type PetrolPrice[XXX]
	Prices []P `bson:"petrolPrices" json:"petrolPrices"`
}

func PlaceResultWPricesArrayByCountry(countryCode string) *any {
	var a any
	switch countryCode {
	case "ar":
		a = []PlaceResultWPrices[petrolprices.PetrolPriceArgentina]{}
	default:
		panic("No case implemented for the countryCode " + countryCode)
	}
	return &a
}
