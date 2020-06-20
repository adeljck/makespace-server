package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"makespace-remaster/module"
	"time"
)

type Project struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id"`
	StartTime   time.Time          `bson:"start_time" json:"start_time"`
	StopTime    time.Time          `bson:"stop_time" json:"stop_time"`
	ProjectName string             `bson:"project_name" json:"project_name"`
	Status      int                `bson:"status" json:"status"`
	ShortInfo   string             `bson:"short_info" json:"short_info"`
	Money       float64            `bson:"money" json:"money"`
	Avatar      string             `bson:"avatar" json:"avatar"`
}

func ProjectList(page int64) (datas []Project, err error) {
	cur, err := module.CLIENT.Mongo.Database("makespace").Collection("projects").Find(context.TODO(), bson.M{}, options.Find().SetSort(bson.M{"create_time": -1}), options.Find().SetLimit(10), options.Find().SetSkip(10*page))
	if err != nil {
		return nil, err
	}
	err = cur.All(context.TODO(), &datas)
	for _, v := range datas {
		if time.Now().Local().After(v.StopTime) && v.Status == 0 {
			_, err = module.CLIENT.Mongo.Database("makespace").Collection("projects").UpdateOne(context.TODO(), bson.M{"_id": v.Id}, bson.D{{"$set", bson.D{{"status", -1}}}})
			v.Status = -1
		}
	}
	return
}
