package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

const collectionName = "gas_stations"

func GetGasStationsForLocation(coord *LatLong, maxDistance *int) *[]PlaceResult {
	filter := bson.D{
		{Key: "locationGEOJson", Value: bson.D{
			{Key: "$near", Value: bson.D{
				{Key: "$geometry", Value: bson.D{
					{Key: "type", Value: "Point"},
					{Key: "coordinates", Value: []float64{coord.Long, coord.Lat}}}},
				{Key: "$maxDistance", Value: *maxDistance}}},
		}},
	}
	cur, err := getMongoClientInstance().Database(dbName).Collection("gas_stations").Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	results := []PlaceResult{}

	if err = cur.All(context.Background(), &results); err != nil {
		panic(err)
	}
	return &results
}
