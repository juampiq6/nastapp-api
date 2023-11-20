package db

type PlaceResult struct {
	Id                    string            `bson:"id" json:"id"`
	DisplayName           string            `bson:"displayName" json:"displayName"`
	AddressComponents     AddressComponents `bson:"addressComponents" json:"addressComponents"`
	FormattedAddress      string            `bson:"formattedAddress" json:"formattedAddress"`
	GoogleMapsUri         string            `bson:"googleMapsUri" json:"googleMapsUri"`
	Location              LatLong           `bson:"location" json:"location"`
	LocationGEOJson       LatLongGEOJson    `bson:"locationGEOJson" json:"locationGEOJson"`
	ShortFormattedAddress string            `bson:"shortFormattedAddress" json:"shortFormattedAddress"`
	WeekdayDescriptions   []string          `bson:"weekdayDescriptions" json:"weekdayDescriptions"`
	PhoneNumber           string            `bson:"nationalPhoneNumber" json:"nationalPhoneNumber"`
	Rating                string            `bson:"rating" json:"rating"`
	UserRatingCount       string            `bson:"userRatingCount" json:"userRatingCount"`
}

type AddressComponents struct {
	Country    string `json:"country" bson:"country"`
	State      string `json:"state" bson:"state"`
	District   string `json:"district" bson:"district"`
	PostalCode string `json:"postalCode" bson:"postalCode"`
}

// Equivalences from the Places Api
const (
	Country    string = "country"
	State             = "administrative_area_level_1"
	District          = "administrative_area_level_2"
	PostalCode        = "postal_code"
)
