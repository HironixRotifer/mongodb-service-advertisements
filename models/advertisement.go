package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Advertisements struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Name        *string            `json:"name" bson:"name" validate:"required,min=2,max=200"`
	Description *string            `json:"description" validate:"required,min=2,max=1000"`
	PhotoLink   *[]string          `json:"photo_link" validate:"required,max=3"`
	Price       *uint64            `json:"price"`
}
