package models

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/meanwhile-app/event-service/configs"
	"github.com/meanwhile-app/event-service/databases"
	"github.com/meanwhile-app/event-service/models/schemas"
	"github.com/meanwhile-app/event-service/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventModel struct {
	Coll *mongo.Collection
}

func NewEventModel() *EventModel {
	env := configs.GetEnv()
	return &EventModel{
		Coll: databases.GetMongoDbClient().Database(env["DB_DATABASE"]).Collection("events"),
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
	if !payload.ReplyToEventId.IsZero() {
		findByIdQuery := bson.M{
			"_id": payload.ReplyToEventId,
		}
		result := eventModel.Coll.FindOne(context.TODO(), findByIdQuery)
		if result.Err() == mongo.ErrNoDocuments {
			return nil, errors.New("invalid reply_to_event_id")
		}
	}

	event := schemas.Event{
		ID:             primitive.NewObjectID(),
		Title:          payload.Title,
		Location:       payload.Location,
		CreatedAt:      time.Now(),
		CreatedBy:      payload.CreatedBy,
		ReplyToEventId: payload.ReplyToEventId,
	}
	result, err := eventModel.Coll.InsertOne(context.TODO(), event)
	if err != nil {
		return nil, err
	}
	return result, nil
}
