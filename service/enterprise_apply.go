package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"makespace-remaster/module"
	"makespace-remaster/serializer"
	"time"
)

type Enterprise struct {
	Name        string    `json:"name" validate:"required,min=2,max=25"`
	Website     string    `json:"site" validate:"required,min=8,max=40"`
	Legal       string    `json:"legal" validate:"required,min=2,max=6"`
	LegalId     string    `json:"legalid" validate:"required,min=18,max=18"`
	Phone       string    `json:"phone" validate:"required,min=5,max=15"`
	Email       string    `json:"email" validate:"required,email,min=8,max=50"`
	Industry    string    `json:"industry" validate:"required,min=2,max=8"`
	Info        string    `json:"info" validate:"required,max=150"`
	Date        time.Time `json:"date" validate:"required"`
	Companycode string    `json:"companycode" validate:"required,min=18,max=18"`
}

func (service *Enterprise) countData(filter interface{}) bool {
	count, _ := module.CLIENT.Mongo.Database("makespace").Collection("enterprise").CountDocuments(context.TODO(), filter)
	if count == 0 {
		return false
	}
	return true
}

// Valid 验证表单
func (service *Enterprise) valid() *serializer.Response {
	var errors []serializer.TagError = make([]serializer.TagError, 0)
	if service.countData(bson.M{"company": service.Name}) {
		errors = append(errors, serializer.TagError{
			Tag:   "name",
			Error: "该公司已入驻",
		})
	}
	if service.countData(bson.M{"register_id": service.LegalId}) {
		errors = append(errors, serializer.TagError{
			Tag:   "legalid",
			Error: "法人身份证已被注册",
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
	if service.countData(bson.M{"companycode": service.Companycode}) {
		errors = append(errors, serializer.TagError{
			Tag:   "companycode",
			Error: "该代码已被注册",
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
func (user *Enterprise) trans() module.Enterprise {
	data := module.Enterprise{
		Phone:       user.Phone,
		Email:       user.Email,
		Name:        user.Name,
		Industry:    user.Industry,
		LegalId:     user.LegalId,
		Date:        user.Date,
		CreateTime:  time.Now(),
		Status:      module.DEACTIVE,
		Avatar:      "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif?imageView2/1/w/80/h/80",
		Legal:       user.Legal,
		Info:        user.Info,
		Website:     user.Website,
		Companycode: user.Companycode,
	}
	return data
}
func (user *Enterprise) Apply(role int) (*serializer.PureErrorResponse, *serializer.Response) {
	if role==1{
		return nil, &serializer.Response{
			Status: 501,
			Data:   nil,
			Msg:    "用户无权限",
		}
	}
	if err := TagValid(user); err != nil {
		return nil, err
	}
	if err := user.valid(); err != nil {
		return nil, err
	}

	module.CLIENT.Mongo.Database("makespace").Collection("enterprise").InsertOne(context.TODO(), user.trans())
	return &serializer.PureErrorResponse{
		Status: 200,
		Msg:    "success",
	}, nil
}
