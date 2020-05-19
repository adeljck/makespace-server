package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"makespace-remaster/module"
	"makespace-remaster/serializer"
	"net/http"
)
// @Summary 学院列表
// @Description 学院列表
// @Produce  json
// @Success 200 {object} serializer.Response "{"status": 200,"data":{"item":[],total:int},"msg": "success" }"
// @Router /academy/list/{page} [get]
func GetAcademy(c *gin.Context) {
	var academies []module.Academy
	cur, err := module.CLIENT.Mongo.Database("makespace").Collection("academy").Find(context.TODO(), bson.M{}, options.Find().SetSort(bson.M{"_id": -1}))
	if err != nil {
		c.JSON(http.StatusOK,ErrorResponse(err))
	}
	err = cur.All(context.TODO(), &academies)
	if err != nil {
		c.JSON(http.StatusOK,ErrorResponse(err))
	}
	res := serializer.BuildListResponse(academies, int64(len(academies)))
	res.Status = http.StatusOK
	res.Msg = "success"
	c.JSON(http.StatusOK, res)
}
// @Summary 行业列表
// @Description 行业列表
// @Produce  json
// @Success 200 {object} serializer.Response "{"status": 200,"data":{"item":[],total:int},"msg": "success" }"
// @Router /industry/list/{page} [get]
func GetIndustry(c *gin.Context) {
	var industries []module.Industry
	cur, err := module.CLIENT.Mongo.Database("makespace").Collection("industry").Find(context.TODO(), bson.M{}, options.Find().SetSort(bson.M{"_id": -1}))
	if err != nil {
		c.JSON(http.StatusOK,ErrorResponse(err))
	}
	err = cur.All(context.TODO(), &industries)
	if err != nil {
		c.JSON(http.StatusOK,ErrorResponse(err))
	}
	res := serializer.BuildListResponse(industries, int64(len(industries)))
	res.Status = http.StatusOK
	res.Msg = "success"
	c.JSON(http.StatusOK, res)
}

