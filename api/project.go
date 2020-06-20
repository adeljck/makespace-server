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

// @Summary 项目列表
// @Description 项目列表
// @Produce  json
// @Param page path int64 true "page"
// @Success 200 {object} serializer.Response "{"status": 200,"data":{"item":[],total:int},"msg": "success" }"
// @Router /projects/{page} [get]
func ProjectList(c *gin.Context) {
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
	if data, err := service.ProjectList(int64(page)); err != nil {
		c.JSON(http.StatusInternalServerError, serializer.PureErrorResponse{
			Status: http.StatusInternalServerError,
			Msg:    err.Error(),
		})
	} else {
		total, _ := module.CLIENT.Mongo.Database("makespace").Collection("projects").CountDocuments(context.Background(), bson.M{})
		res := serializer.BuildListResponse(data, total)
		res.Status = http.StatusOK
		res.Msg = "success"
		c.JSON(200, res)
	}

}

// @Summary 添加项目
// @Description 添加项目
// @accept json
// @Produce  json
// @Param projects body service.ProjectAdd true "project"
// @Success 200 {object} serializer.Response "{"status": 200,"data":bull,"msg": "success" }"
// @Failure 4001 {object}  serializer.Response {"code":4001,"data":null,"msg":"JSON类型不匹配"}
// @Failure 4002 {object}  serializer.Response {"code":4002,"data":null,"msg":"参数错误"}
// @Failure 4003 {object}  serializer.Response {"code":4003,"data":null,"msg":"Tag Error"}
// @Router /project [post]
func AddProject(c *gin.Context) {
	var service service.ProjectAdd
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
// @Summary 删除项目
// @Description 删除项目
// @accept json
// @Produce  json
// @Param projects body service.ProjectDelete true "project"
// @Success 200 {object} serializer.Response "{"status": 200,"data":bull,"msg": "success" }"
// @Failure 4001 {object}  serializer.Response {"code":4001,"data":null,"msg":"JSON类型不匹配"}
// @Failure 4002 {object}  serializer.Response {"code":4002,"data":null,"msg":"参数错误"}
// @Failure 4003 {object}  serializer.Response {"code":4003,"data":null,"msg":"failed"}
// @Router /project [delete]
func DeleteProject(c *gin.Context) {
	var service service.ProjectDelete
	if err := c.ShouldBind(&service); err == nil {
		token := c.Request.Header.Get("Authorization")
		j := middleware.JWT{}
		claims,_ := j.ParseToken(token)
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
// @Summary 接手项目
// @Description 接手项目
// @accept json
// @Produce  json
// @Param projects body service.GetProject true "project"
// @Success 200 {object} serializer.Response "{"status": 200,"data":bull,"msg": "success" }"
// @Failure 4001 {object}  serializer.Response {"code":4001,"data":null,"msg":"JSON类型不匹配"}
// @Failure 4002 {object}  serializer.Response {"code":4002,"data":null,"msg":"参数错误"}
// @Failure 4003 {object}  serializer.Response {"code":4003,"data":null,"msg":"failed"}
// @Router /project [put]
func GetProject(c *gin.Context){
	var service service.GetProject
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
// @Summary 获取项目详细信息
// @Description 获取项目详细信息
// @accept json
// @Param id path string true "id"
// @Produce  json
// @Success 200 {object} serializer.Response "{"status": 200,"data":items,"msg": "success" }"
// @Failure 4001 {object}  serializer.Response {"code":4001,"data":null,"msg":"JSON类型不匹配"}
// @Failure 4002 {object}  serializer.Response {"code":4002,"data":null,"msg":"参数错误"}
// @Failure 4003 {object}  serializer.Response {"code":4003,"data":null,"msg":"failed"}
// @Router /project/{id} [get]
func GetProjectInfo(c *gin.Context){
	id := c.Query("id")
	if res,err := service.GetProjectInfo(id);err != nil{
		c.JSON(http.StatusOK,err)
	}else{
		c.JSON(http.StatusOK,res)
	}
}
