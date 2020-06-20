package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"makespace-remaster/module"
	"makespace-remaster/serializer"
)

type EnterpriseForgetChangePassword struct {
	Password   string ` json:"password" validate:"required,min=8,max=40"`
	Repassword string ` json:"repassword" validate:"required,min=8,max=40"`
}

func (service *EnterpriseForgetChangePassword) valid() *serializer.Response {
	var errors []serializer.TagError = make([]serializer.TagError, 0)
	if service.Repassword != service.Password {
		errors = append(errors, serializer.TagError{
			Tag:   "password",
			Error: "两次输入的密码不相同",
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
func (service *EnterpriseForgetChangePassword) Change(id string) (*serializer.PureErrorResponse, *serializer.Response) {
	if err := TagValid(service); err != nil {
		return nil, err
	}
	if err := service.valid(); err != nil {
		return nil, err
	}
	oid, _ := primitive.ObjectIDFromHex(id)
	module.CLIENT.Mongo.Database("makespace").Collection("user").UpdateOne(context.TODO(), bson.M{"_id": oid}, bson.D{{"$set",bson.D{{"password",module.SetPassword(service.Password)}}}})
	return &serializer.PureErrorResponse{
		Status: 200,
		Msg:    "success",
	}, nil
}
