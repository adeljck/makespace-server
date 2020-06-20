package amdin

import (
	"context"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	"makespace-remaster/module"
	"makespace-remaster/serializer"
	"makespace-remaster/service"
)

// UserLoginService 管理用户登录的服务
type AdminLoginService struct {
	UserName string ` json:"username" validate:"required"`
	Password string ` json:"password" validate:"required"`
}

func (userLoginService *AdminLoginService) valid() *serializer.Response {
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
func (admin *AdminLoginService) countData(filter interface{}) (exists bool) {
	count, _ := module.CLIENT.Mongo.Database("makespace").Collection("user").CountDocuments(context.TODO(), filter)
	if count == 0 {
		return false
	}
	return true
}
func (admin *AdminLoginService) Login() (*service.Info, *serializer.Response) {
	if err := admin.valid(); err != nil {
		return nil, err
	}
	if admin.countData(bson.M{"name": admin.UserName}) == false {
		return nil, &serializer.Response{
			Status: 40005,
			Msg:    "账号或密码错误",
		}
	}

	if module.AdminCheckPassword(admin.UserName, admin.Password) == false {
		return nil, &serializer.Response{
			Status: 40005,
			Msg:    "账号或密码错误",
		}
	}
	var data service.Info
	module.CLIENT.Mongo.Database("makespace").Collection("user").FindOne(context.TODO(), bson.M{"name": admin.UserName}, options.FindOne().SetProjection(bson.D{{"_id", 1}, {"status", 1}, {"name", 1}, {"role", 1}})).Decode(&data)
	return &data, nil
}
