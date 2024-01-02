package schemas

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Location struct {
	Type        string    `json:"type" bson:"type"`               // "Point" for GeoJSON
	Coordinates []float64 `json:"coordinates" bson:"coordinates"` // [longitude, latitude]
}

type Event struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Title     string             `json:"title"`
	Location  Location           `json:"location" bson:"location"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}