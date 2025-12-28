package models

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UUID      uuid.UUID          `bson:"uuid" json:"uuid"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

func NewBaseModel() BaseModel {
	now := time.Now()

	return BaseModel{
		ID:        primitive.NewObjectID(),
		UUID:      uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (b *BaseModel) Touch() {
	b.UpdatedAt = time.Now()
}
