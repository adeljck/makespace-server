package service

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"makespace-remaster/module"
	"makespace-remaster/serializer"
	"time"
	"context"
)

type getProjectInfo struct {
	StartTime    time.Time       `bson:"start_time" json:"start_time"`
	StopTime     time.Time       `bson:"stop_time" json:"stop_time"`
	CreateTime   time.Time       `bson:"create_time" json:"create_time"`
	ProjectName  string          `bson:"project_name" json:"project_name"`
	Creator      string          `bson:"creator" json:"creator"`
	Status       int             `bson:"status" json:"status"`
	Info         string          `bson:"info" json:"info"`
	Attach       []module.Attach `bson:"attach" json:"attach"`
	Image        []module.Image  `bson:"image" json:"image"`
	Money        float64         `bson:"money" json:"money"`
	Contact      module.Contact  `bson:"contact" json:"contact"`
	CommentCount int             `bson:"comment_count" json:"comment_count"`
	Projecter    string          `bson:"projecter" json:"projecter"`
}

func GetProjectInfo(id string) (*serializer.Response, *serializer.PureErrorResponse) {
	var project getProjectInfo
	oid, _ := primitive.ObjectIDFromHex(id)
	module.CLIENT.Mongo.Database("makespace").Collection("projects").FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&project)
	oid,_ = primitive.ObjectIDFromHex(project.Creator)
	type guser struct {
		name string `json:"name", bson:"name"`
	}
	var user guser
	module.CLIENT.Mongo.Database("makespace").Collection("user").FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&user)
	project.Creator = user.name
	//if result != nil {
	//	return nil, &serializer.PureErrorResponse{
	//		Status: 4006,
	//		Msg:    "获取失败",
	//	}
	//} else {
		return &serializer.Response{
			Status: 200,
			Data:   project,
			Msg:    "success",
		}, nil
	//}
}
