package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"makespace-remaster/module"
	"time"
)

type Activity struct {
	Id               primitive.ObjectID `json:"_id" bson:"_id"`
	StartTime        time.Time          `bson:"start_time" json:"start_time"`
	StopTime         time.Time          `bson:"stop_time" json:"stop_time"`
	ActivityName     string             `bson:"activity_name" json:"activity_name"`
	Status           int                `bson:"status" json:"status"`
	ShortInfo        string             `bson:"short_info" json:"short_info"`
	Avatar           string             `bson:"avatar" json:"avatar"`
}

func ActivityList(page int64) (datas []Activity, err error) {
	cur, err := module.CLIENT.Mongo.Database("makespace").Collection("activity").Find(context.TODO(), bson.M{}, options.Find().SetSort(bson.M{"create_time": 1}), options.Find().SetLimit(10), options.Find().SetSkip(10*page))
	if err != nil {
		return nil, err
	}
	err = cur.All(context.Background(), &datas)
	cur.Close(context.Background())
	return
}
