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
	}
	if data,err := service.ProjectList(int64(page));err != nil{
		c.JSON(http.StatusInternalServerError, serializer.PureErrorResponse{
			Status: http.StatusInternalServerError,
			Msg:    err.Error(),
		})
	}else{
		res := serializer.BuildListResponse(data,len(data))
		res.Status = http.StatusOK
		total, _ := module.CLIENT.Mongo.Database("makespace").Collection("projects").CountDocuments(context.Background(), bson.M{})
		res.Msg = total
		c.JSON(200,res)
	}

}
