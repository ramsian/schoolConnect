package healthModel

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HealthModel struct {
	ID          primitive.ObjectID `bson:"_id"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
}

const (
	Name = "health"
)
