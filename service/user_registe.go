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
func (user *User) trans() module.User {
	data := module.User{
		Password:   module.SetPassword(user.Password),
		Phone:      user.Phone,
		Email:      user.Email,
		Name:       user.Name,
		Academy:    user.Academy,
		StudentId:  user.Sid,
		Major:      user.Major,
		Class:      user.Class,
		Date:       user.Date,
		CreateTime: time.Now(),
		Role:       1,
		Avatar:     "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif?imageView2/1/w/80/h/80",
		Status:     module.ACTIVE,
	}
	return data
}
func (service *User) countData(filter interface{}) (exists bool) {
	count, _ := module.CLIENT.Mongo.Database("makespace").Collection("user").CountDocuments(context.TODO(), filter)
	if count == 0 {
		return false
	}
	return true
}
func (service *User) valid() *serializer.Response {
	var errors []serializer.TagError = make([]serializer.TagError, 0)
	if service.Repassword != service.Password {
		errors = append(errors, serializer.TagError{
			Tag:   "password",
			Error: "两次输入的密码不相同",
		})
	}
	if service.countData(bson.M{"name": service.Name, "role": module.NORMAL}) {
		errors = append(errors, serializer.TagError{
			Tag:   "name",
			Error: "该主体已注册账号",
		})
	}
	if service.countData(bson.M{"student_id": service.Sid, "role": module.NORMAL}) {
		errors = append(errors, serializer.TagError{
			Tag:   "sid",
			Error: "学号已被注册",
		})
	}
	if service.countData(bson.M{"email": service.Email}) {
		errors = append(errors, serializer.TagError{
			Tag:   "email",
			Error: "邮箱已被注册",
		})
	}
	if count, _ := module.CLIENT.Mongo.Database("makespace").Collection("academy").CountDocuments(context.TODO(), bson.M{"name": service.Academy});count==0{
		errors = append(errors, serializer.TagError{
			Tag:   "academy",
			Error: "学院不存在",
		})
	}
	if service.countData(bson.M{"phone": service.Phone}) {
		errors = append(errors, serializer.TagError{
			Tag:   "phone",
			Error: "手机号已被注册",
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

func (user *User) Registe() (*serializer.PureErrorResponse, *serializer.Response) {
	if err := TagValid(user); err != nil {
		return nil, err
	}
	if err := user.valid(); err != nil {
		return nil, err
	}
	module.CLIENT.Mongo.Database("makespace").Collection("user").InsertOne(context.TODO(), user.trans())
	return &serializer.PureErrorResponse{
		Status: 200,
		Msg:    "success",
	}, nil
}
