package models

import (
	"context"
	"os"

	"github.com/nuntjw/go-gin-starter/database"
	"github.com/nuntjw/go-gin-starter/models/schemas"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventModel struct {
	Coll *mongo.Collection
}

func NewEventModel() *EventModel {
	return &EventModel{
		Coll: database.GetClient().Database(os.Getenv("DB_DATABASE")).Collection("events"),
	}
}

func (eventModel *EventModel) GetEvents() ([]schemas.Event, error) {
	events := []schemas.Event{}
	cursor, err := eventModel.Coll.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next((context.TODO())) {
		var event schemas.Event
		err := cursor.Decode(&event)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}
