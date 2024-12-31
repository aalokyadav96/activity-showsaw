package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Activity struct {
	Username     string              `json:"username,omitempty" bson:"username,omitempty"`
	PlaceID      string              `json:"placeId,omitempty" bson:"placeId,omitempty"`
	Action       string              `json:"action,omitempty" bson:"action,omitempty"`
	PerformedBy  string              `json:"performedBy,omitempty" bson:"performedBy,omitempty"`
	Timestamp    time.Time           `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
	Details      string              `json:"details,omitempty" bson:"details,omitempty"`
	IPAddress    string              `json:"ipAddress,omitempty" bson:"ipAddress,omitempty"`
	DeviceInfo   string              `json:"deviceInfo,omitempty" bson:"deviceInfo,omitempty"`
	ID           primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	UserID       primitive.ObjectID  `json:"user_id" bson:"user_id"`
	ActivityType string              `json:"activity_type" bson:"activity_type"` // e.g., "follow", "review", "buy"
	EntityID     *primitive.ObjectID `json:"entity_id,omitempty" bson:"entity_id,omitempty"`
	EntityType   *string             `json:"entity_type,omitempty" bson:"entity_type,omitempty"` // "event", "place", or null
}
