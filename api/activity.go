package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"makespace-remaster/middleware"
	"makespace-remaster/module"
	"makespace-remaster/serializer"
	"makespace-remaster/service"
	"net/http"
	"strconv"
	"strings"
)

func ActivityList(c *gin.Context) {
	var page int64
	if c.Param("page") != "/" {
		p, err := strconv.Atoi(strings.Trim(c.Param("page"), "/"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, serializer.PureErrorResponse{
				Status: http.StatusInternalServerError,
				Msg:    err.Error(),
			})
		}
		page = int64(p)
	} else if c.Param("page") == "0" {
		page = 1
	} else {
		page = 1
	}
	if data, err := service.ActivityList(int64(page)); err != nil {
		c.JSON(http.StatusInternalServerError, serializer.PureErrorResponse{
			Status: http.StatusInternalServerError,
			Msg:    err.Error(),
		})
	} else {
		total, _ := module.CLIENT.Mongo.Database("makespace").Collection("activity").CountDocuments(context.Background(), bson.M{})
		res := serializer.BuildListResponse(data, total)
		res.Status = http.StatusOK
		res.Msg = "success"
		c.JSON(200, res)
	}

}
func GetActivityInfo(c *gin.Context) {
	id := c.Query("id")
	if res, err := service.GetActivityInfo(id); err != nil {
		c.JSON(http.StatusOK, err)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func AddActivity(c *gin.Context) {
	var service service.ActivityAdd
	if err := c.ShouldBindJSON(&service); err == nil {
		if success, err := service.Apply(); err == nil {
			c.JSON(http.StatusOK, success)
		} else {
			c.JSON(http.StatusOK, err)
		}
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func DeleteActivity(c *gin.Context) {
	var service service.ActivityDelete
	if err := c.ShouldBind(&service); err == nil {
		token := c.Request.Header.Get("Authorization")
		j := middleware.JWT{}
		claims, _ := j.ParseToken(token)
		if claims.Role !=0{
			c.JSON(http.StatusOK,serializer.PureErrorResponse{
				Status: 500,
				Msg:    "access dined",
			})
		}
		if success, err := service.Delete(claims.Id); err != nil {
			c.JSON(http.StatusOK, err)
		} else {
			c.JSON(http.StatusOK, success)
		}
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func GetActivity(c *gin.Context){
	var service service.GetActivity
	if err := c.ShouldBind(&service); err == nil {
		token := c.Request.Header.Get("Authorization")
		j := middleware.JWT{}
		claims,_ := j.ParseToken(token)
		if success, err := service.Get(claims.Id); err != nil {
			c.JSON(http.StatusOK, err)
		} else {
			c.JSON(http.StatusOK, success)
		}
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}
