package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"makespace-remaster/module"
	"makespace-remaster/serializer"
	"time"
)

type User struct {
	Password   string    `form:"password" json:"password" validate:"required,min=8,max=40"`
	Repassword string    `form:"repassword" json:"repassword" validate:"required,min=8,max=40"`
	Phone      string    `form:"phone" json:"phone" validate:"required,min=5,max=15"`
	Email      string    `form:"email" json:"email" validate:"required,email,min=8,max=50"`
	Name       string    `form:"name" json:"name" validate:"required,min=2,max=6"`
	Academy    string    `form:"academy" json:"academy" validate:"required,min=3,max=20"`
	Sid        string    `form:"sid" json:"sid" validate:"required,min=8,max=15"`
	Major      string    `form:"major" json:"major" validate:"required,min=4,max=15"`
	Class      string    `form:"class" json:"class" validate:"required,min=4,max=10"`
	Date       time.Time `form:"date" json:"date" validate:"required"`
}

func (service *User) valid() *serializer.Response {
	var errors []serializer.TagError = make([]serializer.TagError, 0)
	if service.Repassword != service.Password {
		errors = append(errors, serializer.TagError{
			Tag:   "password",
			Error: "两次输入的密码不相同",
		})
	}
	if CountData(bson.M{"name": service.Name, "role": module.NORMAL}) {
		errors = append(errors, serializer.TagError{
			Tag:   "name",
			Error: "该主体已注册账号",
		})
	}
	if CountData(bson.M{"student_id": service.Sid, "role": module.NORMAL}) {
		errors = append(errors, serializer.TagError{
			Tag:   "sid",
			Error: "学号已被注册",
		})
	}
	if CountData(bson.M{"email": service.Email}) {
		errors = append(errors, serializer.TagError{
			Tag:   "email",
			Error: "邮箱已被注册",
		})
	}
	if CountData(bson.M{"phone": service.Phone}) {
		errors = append(errors, serializer.TagError{
			Tag:   "phone",
			Error: "手机号已被注册",
		})
	}
	if len(errors) != 0 {
		return &serializer.Response{
			Status: 40001,
			Data:   errors,
			Msg:    "something wrong",
		}
	}
	return nil
}

func (user *User) Registe() (*serializer.Response) {
	if err := TagValid(user); err != nil {
		return  err
	}
	if err := user.valid(); err != nil {
		return  err
	}
	module.CLIENT.Mongo.Database("makespace").Collection("user").InsertOne(context.TODO(), Trans(user))
	return nil
}
