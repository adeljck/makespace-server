package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"makespace-remaster/module"
	"makespace-remaster/serializer"
	"time"
)

type ProjectAdd struct {
	StartTime   time.Time       `json:"start_time" validate:"required"`
	StopTime    time.Time       `json:"stop_time" validate:"required"`
	ProjectName string          `json:"project_name" validate:"required,max=40"`
	Creator     string          `json:"creator" validate:"required"`
	Info        string          `json:"info" validate:"required,min=8,max=400"`
	ShortInfo   string          `json:"short_info" validate:"required,min=8,max=200"`
	Attach      []module.Attach `json:"attach" validate:"required"`
	Money       float64         `json:"money" validate:"required"`
	Contact     module.Contact        `json:"contact" validate:"required"`
	Avatar      string          `json:"avatar" validate:"required"`
	Image       []module.Image  `json:"images" validate:"required"`
	ProjectType []string        `json:"project_type" validate:"required"`
}

func (service *ProjectAdd) countData(filter interface{}) bool {
	count, _ := module.CLIENT.Mongo.Database("makespace").Collection("projects").CountDocuments(context.TODO(), filter)
	if count == 0 {
		return false
	}
	return true
}

func (service *ProjectAdd) valid() *serializer.Response {
	var errors []serializer.TagError = make([]serializer.TagError, 0)
	if service.countData(bson.M{"project_name": service.ProjectName}) {
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
		StartTime:   project.StartTime,
		StopTime:    project.StopTime,
		CreateTime:  time.Now(),
		ProjectName: project.ProjectName,
		Creator:     project.Creator,
		Status:      0,
		Info:        project.Info,
		ShortInfo:   project.ShortInfo,
		Attach:      project.Attach,
		Avatar:      project.Avatar,
		Money:       project.Money,
		Image:       project.Image,
		Contact: module.Contact{
			Email:  project.Contact.Email,
			Wechat: project.Contact.Wechat,
			Phone:  project.Contact.Phone,
			Qq:     project.Contact.Qq,
		},
		CommentCount: 0,
		ProjectType:  project.ProjectType,
		Projecter:    "",
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
	_,err := module.CLIENT.Mongo.Database("makespace").Collection("projects").InsertOne(context.TODO(), project.trans())
	if err != nil{
		return nil,&serializer.Response{
			Status: 4003,
			Data:   err,
			Msg:    "failed",
		}
	}
	return &serializer.PureErrorResponse{
		Status: 200,
		Msg:    "success",
	}, nil
}
