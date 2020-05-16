package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"makespace-remaster/module"
	"makespace-remaster/serializer"
	"time"
)

type ProjectAdd struct {
	StartTime   time.Time `json:"start_time" validate:"required"`
	StopTime    time.Time `json:"stop_time" validate:"required"`
	ProjectName string    `json:"project_name" validate:"required,min=8,max=40"`
	Creator     string    `json:"creator" validate:"required,min=8,max=20"`
	Status      int       `json:"status" validate:"required"`
	Info        string    `json:"info" validate:"required,min=8,max=400"`
	ShortInfo   string    `json:"short_info" validate:"required,min=8,max=200"`
	Attach      string    `json:"attach" validate:"required,min=8,max=200"`
	Money       float64   `json:"money" validate:"required"`
	Contact     contact   `json:"contact" validate:"required"`
}
type contact struct {
	Email  string ` json:"email" validate:"required"`
	Wechat string ` json:"wechat" validate:"min=8,max=20"`
	Phone  string ` json:"phone" validate:"min=5,max=11`
	Qq     string ` json:"qq" validate:"min=5,max=13`
}

func (service *ProjectAdd) countData(filter interface{}) bool {
	count, _ := module.CLIENT.Mongo.Database("makespace").Collection("enterprise").CountDocuments(context.TODO(), filter)
	if count == 0 {
		return false
	}
	return true
}

func (service *ProjectAdd) valid() *serializer.Response {
	var errors []serializer.TagError = make([]serializer.TagError, 0)
	if service.countData(bson.M{"company": service.ProjectName}) {
		errors = append(errors, serializer.TagError{
			Tag:   "project_name",
			Error: "该项目已存在",
		})
	}
	if count, _ := module.CLIENT.Mongo.Database("makespace").Collection("industry").CountDocuments(context.TODO(), bson.M{"name": service.ProjectName}); count == 0 {
		errors = append(errors, serializer.TagError{
			Tag:   "project_name",
			Error: "该项目已存在",
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

func (project *ProjectAdd) trans() module.Project {
	data := module.Project{
		StartTime:    project.StartTime,
		StopTime:     project.StopTime,
		CreateTime:   time.Now(),
		ProjectName:  project.ProjectName,
		Creator:      project.Creator,
		Status:       0,
		Info:         project.Info,
		ShortInfo:    project.ShortInfo,
		Attach:       project.Attach,
		Avatar:       project.Attach,
		Money:        project.Money,
		Contact:      module.Contact{
			Email:  project.Contact.Email,
			Wechat: project.Contact.Wechat,
			Phone:  project.Contact.Phone,
			Qq:     project.Contact.Qq,
		},
		CommentCount: 0,
	}
	return data
}
func (project *ProjectAdd) Apply() (*serializer.PureErrorResponse, *serializer.Response) {
	if err := TagValid(project); err != nil {
		return nil, err
	}
	if err := project.valid(); err != nil {
		return nil, err
	}
	module.CLIENT.Mongo.Database("makespace").Collection("projects").InsertOne(context.TODO(), project.trans())
	return &serializer.PureErrorResponse{
		Status: 200,
		Msg:    "success",
	}, nil
}
