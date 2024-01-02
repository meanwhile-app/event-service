package models

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/nuntjw/go-gin-starter/database"
	"github.com/nuntjw/go-gin-starter/models/schemas"
	"github.com/nuntjw/go-gin-starter/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (eventModel *EventModel) GetEvents(filter bson.M) ([]schemas.Event, error) {
	cursor, err := eventModel.Coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	return eventModel.mapCursorToEvents(cursor)
}

func (eventModel *EventModel) GetNearbyEvents(location []string) ([]schemas.Event, error) {
	lat, err := strconv.ParseFloat(location[0], 64)
	if err != nil {
		return nil, err
	}

	lng, err := strconv.ParseFloat(location[1], 64)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"location": bson.M{
			"$nearSphere": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": bson.A{lat, lng},
				},
				"$maxDistance": 100,
			},
		},
	}

	cursor, err := eventModel.Coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	return eventModel.mapCursorToEvents(cursor)
}

func (eventModel *EventModel) mapCursorToEvents(cursor *mongo.Cursor) ([]schemas.Event, error) {
	events := []schemas.Event{}
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

func (eventModel *EventModel) InsertOne(payload *types.InsertEventPayload) (*mongo.InsertOneResult, error) {
	event := schemas.Event{
		ID:        primitive.NewObjectID(),
		Title:     payload.Title,
		Location:  payload.Location,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: payload.CreatedBy,
	}
	result, err := eventModel.Coll.InsertOne(context.TODO(), event)
	if err != nil {
		return nil, err
	}
	return result, nil
}
