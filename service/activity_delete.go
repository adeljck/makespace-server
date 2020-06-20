package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"makespace-remaster/module"
	"makespace-remaster/serializer"
)

type ActivityDelete struct {
	Id string `json:"id" bson:"_id" validate:"required,min=5,max=40"`
}

func (service *ActivityDelete) Delete(id string) (*serializer.Response, *serializer.PureErrorResponse) {
	var user module.User
	oid, err := primitive.ObjectIDFromHex(id)
	module.CLIENT.Mongo.Database("makespace").Collection("user").FindOne(context.TODO(),bson.M{"_id":oid}).Decode(&user)
	if err != nil {
		return nil, &serializer.PureErrorResponse{
			Status: 4003,
			Msg:    "failed",
		}
	}
	if user.Role == 1{
		return nil, &serializer.PureErrorResponse{
			Status: 4003,
			Msg:    "该用户不允许调用",
		}
	}
	oid, err = primitive.ObjectIDFromHex(service.Id)
	if err != nil {
		return nil, &serializer.PureErrorResponse{
			Status: 4003,
			Msg:    "failed",
		}
	}
	_, err = module.CLIENT.Mongo.Database("makespace").Collection("activity").DeleteOne(context.TODO(), bson.M{"_id": oid})
	if err != nil {
		return nil, &serializer.PureErrorResponse{
			Status: 4003,
			Msg:    "failed",
		}
	}
	return &serializer.Response{
		Status: 200,
		Msg:    "success",
	}, nil
}
