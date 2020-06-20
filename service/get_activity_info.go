package service

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"makespace-remaster/module"
	"makespace-remaster/serializer"
	"time"
	"context"
)

type getActivityInfo struct {
	StartTime    time.Time       `bson:"start_time" json:"start_time"`
	StopTime     time.Time       `bson:"stop_time" json:"stop_time"`
	CreateTime   time.Time       `bson:"create_time" json:"create_time"`
	ActivityName string          `bson:"activity_name" json:"activity_name"`
	Creator      string          `bson:"creator" json:"creator"`
	Status       int             `bson:"status" json:"status"`
	Info         string          `bson:"info" json:"info"`
	Attach       []module.Attach `bson:"attach" json:"attach"`
	Avatar       string          `bson:"avatar" json:"avatar"`
	Contact      module.Contact  `bson:"contact" json:"contact"`
	CommentCount int             `bson:"comment_count" json:"comment_count"`
}

func GetActivityInfo(id string) (*serializer.Response, *serializer.PureErrorResponse) {
	var project getActivityInfo
	oid, _ := primitive.ObjectIDFromHex(id)
	result := module.CLIENT.Mongo.Database("makespace").Collection("activity").FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&project)
	oid, _ = primitive.ObjectIDFromHex(project.Creator)
	type guser struct {
		name string `json:"name", bson:"name"`
	}
	var user guser
	module.CLIENT.Mongo.Database("makespace").Collection("user").FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&user)
	project.Creator = user.name
	if result != nil {
		return nil, &serializer.PureErrorResponse{
			Status: 4006,
			Msg:    "获取失败",
		}
	} else {
		return &serializer.Response{
			Status: 200,
			Data:   project,
			Msg:    "success",
		}, nil
	}
}
