package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"makespace-remaster/module"
	"makespace-remaster/serializer"
	"time"
)

type ActivityAdd struct {
	StartTime        time.Time       `json:"start_time" validate:"required"`
	StopTime         time.Time       `json:"stop_time" validate:"required"`
	ActivityName     string          `json:"activity_name" validate:"required,max=40"`
	Creator          string          `json:"creator" validate:"required"`
	Info             string          `json:"info" validate:"required,min=8,max=400"`
	ShortInfo        string          `json:"short_info" validate:"required,min=8,max=200"`
	Attach           []module.Attach `json:"attach" validate:"required"`
	Contact          module.Contact  `json:"contact" validate:"required"`
	Avatar           string          `json:"avatar" validate:"required"`
	ActivityType     string          `json:"activity_type" validate:"required"`
	ActivityLocation string          `json:"activity_location" validate:"required"`
}

func (service *ActivityAdd) countData(filter interface{}) bool {
	count, _ := module.CLIENT.Mongo.Database("makespace").Collection("projects").CountDocuments(context.TODO(), filter)
	if count == 0 {
		return false
	}
	return true
}

func (service *ActivityAdd) valid() *serializer.Response {
	var errors []serializer.TagError = make([]serializer.TagError, 0)
	if service.countData(bson.M{"activity_name": service.ActivityName}) {
		errors = append(errors, serializer.TagError{
			Tag:   "activity_name",
			Error: "该活动名已存在",
		})
	}
	if len(errors) != 0 {
		return &serializer.Response{
			Status: 50001,
			Data:   errors,
			Msg:    "something wrong",
		}
	}
	return nil
}

func (project *ActivityAdd) trans() module.Activity {
	data := module.Activity{
		ActivityLocation: project.ActivityLocation,
		StartTime:        project.StartTime,
		StopTime:         project.StopTime,
		CreateTime:       time.Now(),
		ActivityName:     project.ActivityName,
		Creator:          project.Creator,
		Status:           0,
		Info:             project.Info,
		ShortInfo:        project.ShortInfo,
		Attach:           project.Attach,
		Avatar:           project.Avatar,
		Contact: module.Contact{
			Email:  project.Contact.Email,
			Wechat: project.Contact.Wechat,
			Phone:  project.Contact.Phone,
			Qq:     project.Contact.Qq,
		},
		CommentCount: 0,
		ActivityType: project.ActivityType,
		Joiner:       make([]string, 0),
	}
	return data
}
func (project *ActivityAdd) Apply() (*serializer.PureErrorResponse, *serializer.Response) {
	if err := TagValid(project); err != nil {
		return nil, err
	}
	if err := project.valid(); err != nil {
		return nil, err
	}
	module.CLIENT.Mongo.Database("makespace").Collection("activity").InsertOne(context.TODO(), project.trans())
	return &serializer.PureErrorResponse{
		Status: 200,
		Msg:    "success",
	}, nil
}
