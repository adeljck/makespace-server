package service

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	"log"
	"makespace-remaster/conf"
	"makespace-remaster/module"
	"makespace-remaster/serializer"
	"strconv"
	"time"
)

//字段验证
func TagValid(service interface{}) *serializer.Response {
	validate := validator.New()
	if err := validate.Struct(service); err != nil {
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
			Status: 40002,
			Data:   TagErrors,
			Msg:    "tag error",
		}
	}
	return nil
}
func Trans(user *User) module.User {
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
func CountData(filter interface{}) (exists bool) {
	count, _ := module.CLIENT.Mongo.Database("makespace").Collection("user").CountDocuments(context.TODO(), filter)
	if count == 0 {
		return false
	}
	return true
}

func View(id string) {
	conn, err := conf.RedisPool.Dial()
	if err != nil {
		log.Fatal(err)
	}
	if ok, _ := redis.Bool(conn.Do("EXISTS", id)); ok {
		count, _ := redis.String(conn.Do("GET", id))
		c, _ := strconv.Atoi(count)
		conn.Do("SET", id, c+1)
	} else {
		conn.Do("SET", id, 1)
	}
}
