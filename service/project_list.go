package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"makespace-remaster/module"
	"time"
)

type Project struct {
	StartTime    time.Time      `bson:"start_time" json:"start_time"`
	StopTime     time.Time      `bson:"stop_time" json:"stop_time"`
	ProjectName  string         `bson:"project_name" json:"project_name"`
	Creator      string         `bson:"creator" json:"creator"`
	Status       int            `bson:"status" json:"status"`
	Info         string         `bson:"info" json:"info"`
	ShortInfo    string         `bson:"short_info" json:"short_info"`
	Attach       string         `bson:"attach" json:"attach"`
	Money        string         `bson:"money" json:"money"`
	Contact      module.Contact `bson:"contact" json:"contact"`
	CommentCount int            `bson:"comment_count" json:"comment_count"`
}

func ProjectList(page int64) (data []Project, err error) {
	cur, err := module.CLIENT.Mongo.Database("makespace").Collection("projects").Find(context.TODO(), bson.M{}, options.Find().SetSort(bson.M{"_id": -1}), options.Find().SetLimit(10), options.Find().SetSkip(10*page))
	if err != nil {
		return nil, err
	}
	err = cur.All(context.TODO(), &data)
	return
}
