package db

// @Description LatLong struct
type LatLong struct {
	Lat  float64 `json:"lat" form:"lat" binding:"required"`
	Long float64 `json:"long" form:"long" binding:"required"`
}

// For this program purpose, only point type is used
type LatLongGEOJson struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
}
