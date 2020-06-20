package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"makespace-remaster/module"
	"makespace-remaster/serializer"
	"time"
)

type EnterpriseUser struct {
	Password   string    ` json:"password" validate:"required,min=8,max=40"`
	Repassword string    ` json:"repassword" validate:"required,min=8,max=40"`
	Phone       string    ` json:"phone" validate:"required,min=5,max=15"`
	Email       string    ` json:"email" validate:"required,email,min=8,max=50"`
	Name        string    ` json:"name" validate:"required,min=2,max=6"`
	Id          string    ` json:"id" validate:"required,min=18,max=18"`
	Industry    string    ` json:"industry" validate:"required,min=3,max=10"`
	CompanyCode string    ` json:"company_code" validate:"required,min=18,max=18"`
	Company     string    ` json:"company" validate:"required,min=3,max=15"`
}

func (user *EnterpriseUser) trans() module.EnterpriseUser {
	data := module.EnterpriseUser{
		Password:   module.SetPassword(user.Password),
		Phone:      user.Phone,
		Email:      user.Email,
		Name:       user.Name,
		Id: user.Id,
		Industry: user.Industry,
		Company: user.Company,
		CompanyCode: user.CompanyCode,
		CreateTime: time.Now(),
		Role:       2,
		Avatar:     "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif?imageView2/1/w/80/h/80",
		Status:     module.ACTIVE,
	}
	return data
}
func (service *EnterpriseUser) countData(filter interface{}) (exists bool) {
	count, _ := module.CLIENT.Mongo.Database("makespace").Collection("user").CountDocuments(context.TODO(), filter)
	if count == 0 {
		return false
	}
	return true
}
func (service *EnterpriseUser) valid() *serializer.Response {
	var errors []serializer.TagError = make([]serializer.TagError, 0)
	if service.Repassword != service.Password {
		errors = append(errors, serializer.TagError{
			Tag:   "password",
			Error: "两次输入的密码不相同",
		})
	}
	if service.countData(bson.M{"name": service.Name, "role": 2}) {
		errors = append(errors, serializer.TagError{
			Tag:   "name",
			Error: "该主体已注册账号",
		})
	}
	if service.countData(bson.M{"id": service.Id, "role": 2}) {
		errors = append(errors, serializer.TagError{
			Tag:   "id",
			Error: "身份证号已被注册",
		})
	}
	if service.countData(bson.M{"email": service.Email}) {
		errors = append(errors, serializer.TagError{
			Tag:   "email",
			Error: "邮箱已被注册",
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
func (user *EnterpriseUser) Registe() (*serializer.PureErrorResponse, *serializer.Response) {
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

