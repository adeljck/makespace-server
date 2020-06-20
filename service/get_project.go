package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"makespace-remaster/module"
	"makespace-remaster/serializer"
	"time"
)

type GetProject struct {
	Id string `json:"id" bson:"_id"`
}

func (service *GetProject) Get(id string) (*serializer.Response, *serializer.PureErrorResponse) {
	oid, err := primitive.ObjectIDFromHex(service.Id)
	var project module.Project
	module.CLIENT.Mongo.Database("makespace").Collection("projects").FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&project)
	if project.Status == 1{
		return nil, &serializer.PureErrorResponse{
			Status: 4003,
			Msg:    "已被竞标",
		}
	}
	if project.Status == -1{
		return nil, &serializer.PureErrorResponse{
			Status: 4003,
			Msg:    "已结束",
		}
	}
	if time.Now().After(project.StopTime){
		_, err = module.CLIENT.Mongo.Database("makespace").Collection("projects").UpdateOne(context.TODO(), bson.M{"_id": oid}, bson.D{{"$set", bson.D{{"status",-1}}}})
		return nil, &serializer.PureErrorResponse{
			Status: 4003,
			Msg:    "项目已结束",
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
	_, err = module.CLIENT.Mongo.Database("makespace").Collection("projects").UpdateOne(context.TODO(), bson.M{"_id": oid}, bson.D{{"$set", bson.D{{"projecter", id},{"status",1}}}})
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
