package service

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"makespace-remaster/module"
	"makespace-remaster/serializer"
	"context"
)

type EnterpriseUserForget struct {
	Id   string ` bson:"id" json:"id" validate:"required,min=15,max=18"`
	Phone string ` bson:"phone" json:"phone" validate:"required,min=5,max=15"`
	Email string ` bson:"email" json:"email" validate:"required,email,min=8,max=50"`
}

func (service *EnterpriseUserForget) countData(filter interface{}) (exists bool) {
	count, _ := module.CLIENT.Mongo.Database("makespace").Collection("user").CountDocuments(context.TODO(), filter)
	if count == 0 {
		return true
	}
	return false
}
func (service *EnterpriseUserForget) Forget() (*Info, *serializer.Response) {
	if err := TagValid(service); err != nil {
		return nil, err
	}
	if service.countData(bson.M{"id": service.Id}) {
		return nil, &serializer.Response{
			Status: 40005,
			Msg:    "账号不存在",
		}
	}
	var check EnterpriseUserForget
	module.CLIENT.Mongo.Database("makespace").Collection("user").FindOne(context.TODO(), bson.M{"id": service.Id}, options.FindOne().SetProjection(bson.D{{"_id", 0}})).Decode(&check)
	if service.Email == check.Email && service.Phone == check.Phone {
		var data Info
		module.CLIENT.Mongo.Database("makespace").Collection("user").FindOne(context.TODO(), bson.M{"id": service.Id}).Decode(&data)
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
