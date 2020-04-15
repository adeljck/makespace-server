package module

import "time"

type Bussiness struct {
	UserName   string    `bson:"username" json:"username"`
	Password   string    `bson:"password" json:"password"`
	Phone      string    `bson:"phone" json:"phone"`
	Email      string    `bson:"email" json:"email"`
	Legal      string    `bson:"legal" json:"legal"`
	Info       string    `bson:"info" json:"info"`
	Website    string    `bson:"website" json:"website"`
	Company    string    `bson:"company" json:"company"`
	RegisterId string    `bson:"register_id" json:"register_id"`
	Role       int       `bson:"role" json:"role"`
	Status     int       `bson:"status" json:"status"`
	CreateTime time.Time `bson:"create_time" json:"create_time"`
	Avatar     string    `bson:"avatar" json:"avatar"`
}
