package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Advertisements struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Name        *string            `json:"name" validate:"requred,min=2,max=200"`
	Description *string            `json:"description" validate:"requred,min=2,max=1000"`
	PhotoLink   *[]string          `json:"photo_link" validate:"requred, max=3"`
	Price       *uint64            `json:"price"`
}
