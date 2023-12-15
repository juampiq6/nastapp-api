package db

import "nastapp-api/db/petrolprices"

type PlaceResultWPrices struct {
	MongoId string `bson:"_id" json:"_id"`
	// GoogleMapsId, not useful for frontend
	// Id                string            `bson:"id" json:"id"`
	DisplayName       string            `bson:"displayName" json:"displayName"`
	AddressComponents AddressComponents `bson:"addressComponents" json:"addressComponents"`
	// FormattedAddress  string            `bson:"formattedAddress" json:"formattedAddress"`
	GoogleMapsUri string  `bson:"googleMapsUri" json:"googleMapsUri"`
	Location      LatLong `bson:"location" json:"location"`
	// LocationGEOJson       LatLongGEOJson    `bson:"locationGEOJson" json:"locationGEOJson"`
	ShortFormattedAddress string                              `bson:"shortFormattedAddress" json:"shortFormattedAddress"`
	WeekdayDescriptions   []string                            `bson:"weekdayDescriptions" json:"weekdayDescriptions"`
	PhoneNumber           string                              `bson:"nationalPhoneNumber" json:"nationalPhoneNumber"`
	Rating                string                              `bson:"rating" json:"rating"`
	UserRatingCount       string                              `bson:"userRatingCount" json:"userRatingCount"`
	Prices                []petrolprices.PetrolPriceArgentina `bson:"petrolPrices" json:"petrolPrices"`
}
