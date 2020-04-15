package service
//
//import (
//	"go.mongodb.org/mongo-driver/bson"
//	"makespace-remaster/module"
//	"makespace-remaster/serializer"
//)
//
//type Bussiness struct {
//	UserName        string `form:"username" json:"username" validate:"required,min=5,max=30,alphanum"`
//	Password        string `form:"password" json:"password" validate:"required,min=8,max=40"`
//	PasswordConfirm string `form:"password_confirm" json:"password_confirm" validate:"required,min=8,max=40"`
//	Phone           string `form:"phone" json:"phone" validate:"required,min=5,max=15"`
//	Email           string `form:"email" json:"email" validate:"required,email,min=8,max=50"`
//	Legal           string `form:"legal" json:"legal" validate:"required,min=2,max=6"`
//	Info            string `form:"info" json:"info" validate:"required,max=150"`
//	Website         string `form:"pwebsite" json:"website" validate:"required,min=8,max=40"`
//	Company         string `form:"company" json:"company" validate:"required,min=2,max=25"`
//	RegisterId      string `form:"registerid" json:"registerid" validate:"required,min=18,max=18"`
//}
//
//// Valid 验证表单
//func (service *Bussiness) valid() *serializer.Response {
//	var errors []serializer.TagError = make([]serializer.TagError, 0)
//	if service.PasswordConfirm != service.Password {
//		errors = append(errors, serializer.TagError{
//			Tag:   "password",
//			Error: "两次输入的密码不相同",
//		})
//	}
//	if CountData(bson.M{"username": service.UserName}) {
//		errors = append(errors, serializer.TagError{
//			Tag:   "username",
//			Error: "用户名已被注册",
//		})
//	}
//	if CountData(bson.M{"company": service.Company}) {
//		errors = append(errors, serializer.TagError{
//			Tag:   "Company",
//			Error: "该主体已入驻",
//		})
//	}
//	if CountData(bson.M{"register_id": service.RegisterId}) {
//		errors = append(errors, serializer.TagError{
//			Tag:   "register_id",
//			Error: "身份证已被注册",
//		})
//	}
//	if CountData(bson.M{"email": service.Email}) {
//		errors = append(errors, serializer.TagError{
//			Tag:   "email",
//			Error: "邮箱已被注册",
//		})
//	}
//	if CountData(bson.M{"phone": service.Phone}) {
//		errors = append(errors, serializer.TagError{
//			Tag:   "phone",
//			Error: "手机号已被注册",
//		})
//	}
//	if len(errors) != 0 {
//		return &serializer.Response{
//			Status: 40001,
//			Data:   errors,
//			Msg:    "something wrong",
//		}
//	}
//	return nil
//}
//
//func (user *Bussiness) Registe() (interface{}, *serializer.Response) {
//	if err := TagValid(user); err != nil {
//		return nil, err
//	}
//	if err := user.valid(); err != nil {
//		return nil, err
//	}
//	insertResult, _ := module.CLIENT.Mongo.Database("makespace").Collection("user").InsertOne(context.TODO(), Trans(user))
//	return insertResult.InsertedID, nil
//}
