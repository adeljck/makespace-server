package service

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	"makespace-remaster/serializer"
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
			Status: 40003,
			Data:   TagErrors,
			Msg:    "tag error",
		}
	}
	return nil
}

