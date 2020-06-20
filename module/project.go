package module

import (
	"time"
)

type Contact struct {
	Email  string `bson:"email" json:"email"`
	Wechat string `bson:"wechat" json:"wechat"`
	Phone  string `bson:"phone" json:"phone"`
	Qq     string `bson:"qq" json:"qq"`
}
type Comment struct {
	CommentProject string    `bson:"comment_project" json:"comment_project"`
	Poster         string    `bson:"poster" json:"poster"`
	Content        string    `bson:"Content" json:"Content"`
	IsDeleted      int       `bson:"is_deleted" json:"is_deleted"`
	CreateTime     time.Time `bson:"create_time" json:"create_time"`
}
type Attach struct {
	Name string `bson:"name" json:"name"`
	Url  string `bson:"url" json:"url"`
}
type Image struct {
	Name string `bson:"name" json:"name"`
	Url  string `bson:"url" json:"url"`
}
type Project struct {
	StartTime    time.Time `bson:"start_time" json:"start_time"`
	StopTime     time.Time `bson:"stop_time" json:"stop_time"`
	CreateTime   time.Time `bson:"create_time" json:"create_time"`
	ProjectName  string    `bson:"project_name" json:"project_name"`
	Creator      string    `bson:"creator" json:"creator"`
	Status       int       `bson:"status" json:"status"`
	Info         string    `bson:"info" json:"info"`
	ShortInfo    string    `bson:"short_info" json:"short_info"`
	Attach       []Attach  `bson:"attach" json:"attach"`
	Avatar       string    `bson:"avatar" json:"avatar"`
	Image        []Image   `bson:"image" json:"image"`
	Money        float64   `bson:"money" json:"money"`
	Contact      Contact   `bson:"contact" json:"contact"`
	CommentCount int       `bson:"comment_count" json:"comment_count"`
	ProjectType  []string  `bson:"project_type json:"project_type`
	Projecter    string    `bson:"projecter" json:"projecter"`
}
