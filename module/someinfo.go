package module

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Academy struct {
	Id    primitive.ObjectID `json:"_id" bson:"_id"`
	Name  string             `json:"name" bson:"name"`
	Major []string       `json:"major" bson:"major"`
}
type Industry struct {
	Id   primitive.ObjectID `json:"_id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
}
