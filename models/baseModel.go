package models

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id" swaggertype:"string" readonly:"true" example:"auto-generated"`
	UUID      uuid.UUID          `bson:"uuid" json:"uuid" readonly:"true" example:"auto-generated"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at" readonly:"true" example:"2026-01-13T12:00:00Z"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at" readonly:"true" example:"2026-01-13T12:00:00Z"`
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
