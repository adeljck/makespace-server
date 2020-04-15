package module

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

type Database struct {
	Mongo *mongo.Client
}

var
(
	CLIENT   *Database
	DATABASE string = os.Getenv("MONGO_DATABASE")
	ACTIVE   int    = 1
	DEACTIVE int    = 0
	BAN             = -1
	NORMAL   int    = 1
	ADMIN    int    = 0
	CHARTOR         = map[int]string{
		ADMIN:  "admin",
		NORMAL: "NORMAL",
	}
)

//初始化
func MongoInit() {
	CLIENT = &Database{
		Mongo: SetConnect(),
	}
	err := CLIENT.Mongo.Ping(context.TODO(), readpref.Nearest())
	if err != nil {
		log.Fatal(err)
	}
	if count, _ := CLIENT.Mongo.Database("makespace").Collection("user").CountDocuments(context.TODO(), bson.M{"role": 0}); count == 0 {
		data := User{
			Name:       os.Getenv("ADMIN_NAME"),
			Password:   SetPassword(os.Getenv("ADMIN_PASSWORD")),
			CreateTime: time.Now(),
			Role:       0,
			Avatar:     "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif?imageView2/1/w/80/h/80",
			Date:       time.Now(),
		}
		CLIENT.Mongo.Database("makespace").Collection("user").InsertOne(context.TODO(), data)
	}
}

// 连接设置
func SetConnect() *mongo.Client {
	uri := os.Getenv("MONGO_URI")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetMaxPoolSize(500)) // 连接池
	if err != nil {
		fmt.Println(err)
	}
	return client
}
