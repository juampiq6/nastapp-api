package petrolprices

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PetrolPriceArgentina struct {
	StationID    primitive.ObjectID `json:"stationId,required" bson:"stationId,required"`
	Timestamp    time.Time          `json:"timestamp" bson:"timestamp"`
	PetrolGrade2 float64            `json:"petrolGrade2,omitempty" bson:"petrolGrade2,omitempty"`
	PetrolGrade3 float64            `json:"petrolGrade3,omitempty" bson:"petrolGrade3,omitempty"`
	DieselGrade2 float64            `json:"dieselGrade2,omitempty" bson:"dieselGrade2,omitempty"`
	DieselGrade3 float64            `json:"dieselGrade3,omitempty" bson:"dieselGrade3,omitempty"`
	Cng          float64            `json:"cng,omitempty" bson:"cng,omitempty"`
}
