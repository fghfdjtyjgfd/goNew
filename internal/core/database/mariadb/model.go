package database

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Log Model log model
type LogModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	IP        string             `bson:"ip,omitempty"`
	Method    string             `bson:"method,omitempty"`
	Uri       string             `bson:"uri,omitempty"`
	BodyJson  []byte             `bson:"body_json,omitempty"`
}

// Stamp Model stamp model
func (model *LogModel) Stamp() {
	model.ID = primitive.NewObjectID()
	model.CreatedAt = time.Now()
}
