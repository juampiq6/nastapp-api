package db

import (
	"context"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
)

const stationsCollectionName = "gas_stations"
const petrolPricesCollectionPrefix = "petrol_price_"

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

func GetGasStationsAndPricesForLocation(coord *LatLong, maxDistance *int, countryCode string) *[]PlaceResultWPrices {
	collection := petrolPricesCollectionPrefix + strings.ToLower(countryCode)
	metersToRad := float64(*maxDistance) / float64(6378100)
	filterPhase :=
		bson.D{{Key: "$match", Value: bson.D{
			{Key: "locationGEOJson", Value: bson.D{
				{Key: "$geoWithin", Value: bson.D{
					{Key: "$centerSphere", Value: []interface{}{[]float64{coord.Long, coord.Lat}, metersToRad}},
				}},
			}},
		}},
		}
	joinPhase := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: collection},
			{Key: "localField", Value: "_id"},
			{Key: "foreignField", Value: "stationId"},
			{Key: "as", Value: "petrolPrices"},
			{Key: "pipeline", Value: []bson.D{
				{
					{Key: "$sort", Value: bson.D{{Key: "timestamp", Value: -1}}},
				},
				{
					{Key: "$limit", Value: 3},
				},
			}},
		}},
	}

	cur, err := getMongoClientInstance().Database(dbName).Collection("gas_stations").Aggregate(context.TODO(), []interface{}{filterPhase, joinPhase})
	if err != nil {
		panic(err)
	}
	results := []PlaceResultWPrices{}

	if err = cur.All(context.Background(), &results); err != nil {
		panic(err)
	}
	return &results
}

func InsertGasPriceForStation(prices *any, countryCode string) (string, error) {
	collection := petrolPricesCollectionPrefix + strings.ToLower(countryCode)
	r, err := getMongoClientInstance().Database(dbName).Collection(collection).InsertOne(context.TODO(), *prices)
	if err != nil {
		return "", err
	}
	return (r.InsertedID.(primitive.ObjectID)).Hex(), nil
}
