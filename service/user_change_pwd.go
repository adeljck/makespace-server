package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"makespace-remaster/module"
	"makespace-remaster/serializer"
)

type UserChangePassword struct {
	OldPassword string ` json:"oldpassword" validate:"required,min=8,max=40"`
	Password    string ` json:"password" validate:"required,min=8,max=40"`
	Repassword  string ` json:"repassword" validate:"required,min=8,max=40"`
}
func (service *UserChangePassword) valid(id string) *serializer.Response {
	var errors []serializer.TagError = make([]serializer.TagError, 0)
	var user module.User
	oid ,_ := primitive.ObjectIDFromHex(id)
	module.CLIENT.Mongo.Database("makespace").Collection("user").FindOne(context.TODO(), bson.M{"_id":oid }, options.FindOne().SetProjection(bson.M{"password": 1})).Decode(&user)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(service.OldPassword))
	if err != nil{
		errors = append(errors, serializer.TagError{
			Tag:   "oldpassword",
			Error: "旧密码错误",
		})
	}
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
func (service *UserChangePassword) Change(id string) (*serializer.PureErrorResponse, *serializer.Response) {
	if err := TagValid(service); err != nil {
		return nil, err
	}
	if err := service.valid(id); err != nil {
		return nil, err
	}
	oid, _ := primitive.ObjectIDFromHex(id)
	module.CLIENT.Mongo.Database("makespace").Collection("user").UpdateOne(context.TODO(), bson.M{"_id": oid}, bson.D{{"$set",bson.D{{"password",module.SetPassword(service.Password)}}}})
	return &serializer.PureErrorResponse{
		Status: 200,
		Msg:    "success",
	}, nil
}