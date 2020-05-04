package service

import (
	"context"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	"makespace-remaster/module"
	"makespace-remaster/serializer"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"username" json:"username" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}
type Info struct {
	Name   string             `form:"name" json:"name" bson:"name"`
	Status int             `form:"status" json:"status" bson:"status"`
	ID     primitive.ObjectID `form:"_id" json:"_id" bson:"_id"`
}

func (userLoginService *UserLoginService) valid() *serializer.Response {
	validate := validator.New()
	if err := validate.Struct(userLoginService); err != nil {
		trans, _ := ut.New(zh.New()).GetTranslator("zh")
		zh_translations.RegisterDefaultTranslations(validate, trans)
		var TagErrors []serializer.TagError = make([]serializer.TagError, 0)
		for _, err := range err.(validator.ValidationErrors) {
			tagerror := serializer.TagError{
				Tag:   err.Field(),
				Error: err.Translate(trans),
			}
			TagErrors = append(TagErrors, tagerror)
		}
		return &serializer.Response{
			Status: 40003,
			Data:   TagErrors,
			Msg:    "tag error",
		}
	}
	return nil
}
func (service *UserLoginService) countData(filter interface{}) (exists bool) {
	count, _ := module.CLIENT.Mongo.Database("makespace").Collection("user").CountDocuments(context.TODO(), filter)
	if count == 0 {
		return false
	}
	return true
}
func (service *UserLoginService) Login() (*Info, *serializer.Response) {
	if err := service.valid(); err != nil {
		return nil, err
	}
	if service.countData(bson.M{"email": service.UserName}) {
		return nil, &serializer.Response{
			Status: 40005,
			Msg:    "账号或密码错误",
		}
	}

	if module.CheckPassword(service.UserName, service.Password) == false {
		return nil, &serializer.Response{
			Status: 40005,
			Msg:    "账号或密码错误",
		}
	}
	var data Info
	module.CLIENT.Mongo.Database("makespace").Collection("user").FindOne(context.TODO(), bson.M{"email": service.UserName}, options.FindOne().SetProjection(bson.D{{"_id", 1}, {"status", 1}, {"name", 1}})).Decode(&data)
	if data.Status == module.DEACTIVE {
		return nil, &serializer.Response{
			Status: 40006,
			Data:   nil,
			Msg:    "unauthrized account",
		}
	}
	return &data, nil
}
