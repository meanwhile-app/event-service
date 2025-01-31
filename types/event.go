package types

import (
	"github.com/meanwhile-app/event-service/models/schemas"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InsertEventPayload struct {
	Title          string              `json:"title"`
	Location       schemas.Location    `json:"location" bson:"location"`
	CreatedBy      primitive.ObjectID  `json:"created_by" bson:"created_by"`
	ReplyToEventId *primitive.ObjectID `json:"reply_to_event_id,omitempty" bson:"reply_to_event_id,omitempty"`
}
