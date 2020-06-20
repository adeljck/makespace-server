package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"makespace-remaster/module"
	"makespace-remaster/serializer"
	"time"
)

type GetActivity struct {
	Id string `json:"id" bson:"_id" validate:"required,min=5,max=40"`
}

func (service *GetActivity) Get(id string) (*serializer.Response, *serializer.PureErrorResponse) {
	oid, err := primitive.ObjectIDFromHex(service.Id)
	var project module.Project
	module.CLIENT.Mongo.Database("makespace").Collection("projects").FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&project)
	if project.StopTime.After(time.Now()){
		return nil, &serializer.PureErrorResponse{
			Status: 4003,
			Msg:    "活动已结束",
		}
	}
	var user module.User
	oid, err = primitive.ObjectIDFromHex(id)
	module.CLIENT.Mongo.Database("makespace").Collection("user").FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&user)
	if err != nil {
		return nil, &serializer.PureErrorResponse{
			Status: 4003,
			Msg:    "failed",
		}
	}
	if user.Role == 2 {
		return nil, &serializer.PureErrorResponse{
			Status: 4003,
			Msg:    "该用户不允许调用",
		}
	}
	oid, err = primitive.ObjectIDFromHex(service.Id)
	var activit module.Activity
	module.CLIENT.Mongo.Database("makespace").Collection("activity").FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&activit)
	for _,v := range activit.Joiner{
		if v == id{
			return nil, &serializer.PureErrorResponse{
				Status: 4003,
				Msg:    "你已参加本活动",
			}
		}
	}
	_, err = module.CLIENT.Mongo.Database("makespace").Collection("activity").UpdateOne(context.TODO(), bson.M{"_id": oid}, bson.D{{"$addToSet", bson.D{{"joiner", id}}}})
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
