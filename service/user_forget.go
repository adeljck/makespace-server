package service

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"makespace-remaster/module"
	"makespace-remaster/serializer"
	"context"
)

type UserForget struct {
	Sid   string ` json:"sid" validate:"required,min=8,max=15"`
	Phone string ` json:"phone" validate:"required,min=5,max=15"`
	Email string ` json:"email" validate:"required,email,min=8,max=50"`
}

func (service *UserForget) countData(filter interface{}) (exists bool) {
	count, _ := module.CLIENT.Mongo.Database("makespace").Collection("user").CountDocuments(context.TODO(), filter)
	if count == 0 {
		return true
	}
	return false
}
func (service *UserForget) Forget() (*Info, *serializer.Response) {
	if err := TagValid(service); err != nil {
		return nil, err
	}
	if service.countData(bson.M{"student_id": service.Sid}) {
		return nil, &serializer.Response{
			Status: 40005,
			Msg:    "账号不存在",
		}
	}
	var check UserForget
	module.CLIENT.Mongo.Database("makespace").Collection("user").FindOne(context.TODO(), bson.M{"student_id": service.Sid}, options.FindOne().SetProjection(bson.D{{"_id", 0}, {"email", 1}, {"phone", 1}})).Decode(&check)
	if service.Email == check.Email && service.Phone == check.Phone {
		var data Info
		module.CLIENT.Mongo.Database("makespace").Collection("user").FindOne(context.TODO(), bson.M{"student_id": service.Sid}, options.FindOne().SetProjection(bson.D{{"_id", 1}, {"status", 1}, {"name", 1}})).Decode(&data)
		if data.Status == module.DEACTIVE {
			return nil, &serializer.Response{
				Status: 40006,
				Data:   nil,
				Msg:    "unauthrized account",
			}
		}else if data.Role == 0{
			return nil, &serializer.Response{
				Status: 40006,
				Data:   nil,
				Msg:    "wrong account",
			}
		}
		return &data, nil
	}
	return nil, &serializer.Response{
		Status: 40006,
		Data:   nil,
		Msg:    "信息不匹配",
	}
}
