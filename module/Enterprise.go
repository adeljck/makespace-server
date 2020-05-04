package module

import "time"

type Enterprise struct {
	Name        string    `bson:"name" json:"name"`
	Phone       string    `bson:"phone" json:"phone"`
	Email       string    `bson:"email" json:"email"`
	Legal       string    `bson:"legal" json:"legal"`
	Info        string    `bson:"info" json:"info"`
	Website     string    `bson:"website" json:"website"`
	Companycode string    `bson:"company_code" json:"company_code"`
	Industry    string    `bson:"industry" json:"industry"`
	LegalId     string    `bson:"legal_id" json:"legal_id"`
	Status      int       `bson:"status" json:"status"`
	CreateTime  time.Time `bson:"create_time" json:"create_time"`
	Avatar      string    `bson:"avatar" json:"avatar"`
	Date   time.Time `bson:"build_time" json:"build_time"`
}
