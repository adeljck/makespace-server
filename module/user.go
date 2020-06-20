package module

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Password   string    `bson:"password" json:"password"`
	Phone      string    `bson:"phone" json:"phone"`
	Email      string    `bson:"email" json:"email"`
	Name       string    `bson:"name" json:"name"`
	Academy    string    `bson:"academy" json:"academy"`
	StudentId  string    `bson:"student_id" json:"student_id"`
	Major      string    `bson:"major" json:"major"`
	Class      string    `bson:"class" json:"class"`
	Role       int       `bson:"role" json:"role"`
	Status     int       `bson:"status" json:"status"`
	Date       time.Time `bson:"date" json:"date"`
	CreateTime time.Time `bson:"create_time" json:"create_time"`
	Avatar     string    `bson:"avatar" json:"avatar"`
}
type EnterpriseUser struct {
	Password    string    `bson:"password" json:"password"`
	Phone       string    `bson:"phone" json:"phone"`
	Email       string    `bson:"email" json:"email"`
	Name        string    `bson:"name" json:"name"`
	Id          string    `bson:"id" json:"id"`
	Industry    string    `bson:"industry" json:"industry"`
	Role        int       `bson:"role" json:"role"`
	Status      int       `bson:"status" json:"status"`
	CompanyCode string    `bson:"company_code" json:"company"`
	Company     string    `bson:"company" json:"company"`
	CreateTime  time.Time `bson:"create_time" json:"create_time"`
	Avatar      string    `bson:"avatar" json:"avatar"`
}

// SetPassword 设置密码
func SetPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	passwordDigest := string(bytes)
	return passwordDigest
}

// CheckPassword 校验密码
func CheckPassword(email string, password string) bool {
	var user User
	CLIENT.Mongo.Database("makespace").Collection("user").FindOne(context.TODO(), bson.M{"email": email}, options.FindOne().SetProjection(bson.M{"password": 1})).Decode(&user)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}
// CheckPassword 校验密码
func AdminCheckPassword(email string, password string) bool {
	var user User
	CLIENT.Mongo.Database("makespace").Collection("user").FindOne(context.TODO(), bson.M{"name": email}, options.FindOne().SetProjection(bson.M{"password": 1})).Decode(&user)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}
