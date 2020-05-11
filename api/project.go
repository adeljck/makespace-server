package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"makespace-remaster/module"
	"makespace-remaster/serializer"
	"makespace-remaster/service"
	"net/http"
	"strconv"
	"strings"
)
// @Summary 项目列表
// @Description 项目列表
// @accept json
// @Produce  json
// @Param projects body service.Project true "project"
// @Success 200 {object} serializer.Response "{"status": 200,"data":{"item":[],total:0},"msg": "success" }"
// @Router /project/list/{page} [get]
func ProjectList(c *gin.Context) {
	var page int64
	if c.Param("page") != "/"{
		p, err := strconv.Atoi(strings.Trim(c.Param("page"),"/"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, serializer.PureErrorResponse{
				Status: http.StatusInternalServerError,
				Msg:    err.Error(),
			})
		}
		page = int64(p)
	}else{
		page = 1
	}
	if data,err := service.ProjectList(int64(page));err != nil{
		c.JSON(http.StatusInternalServerError, serializer.PureErrorResponse{
			Status: http.StatusInternalServerError,
			Msg:    err.Error(),
		})
	}else{
		total, _ := module.CLIENT.Mongo.Database("makespace").Collection("projects").CountDocuments(context.Background(), bson.M{})
		res := serializer.BuildListResponse(data,total)
		res.Status = http.StatusOK
		res.Msg = "success"
		c.JSON(200,res)
	}

}
func SearchProject(c *gin.Context){
	c.PostForm("")
}
