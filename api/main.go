package api

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"makespace-remaster/middleware"
	"makespace-remaster/serializer"
	"net/http"
)

func ErrorResponse(err error) serializer.Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: 40001,
			Msg:    "JSON类型不匹配",
		}
	}

	return serializer.Response{
		Status: 40002,
		Msg:    "参数错误",
	}
}
// @Summary token检查
// @Description token检查
// @accept */*
// @Produce  json
// @Success 200 {object} serializer.PureErrorResponse "{"status": 0,"msg": "success" }"
// @Failure -1 {object}  serializer.PureErrorResponse {"code":-1,"msg":"error"}
// @Failure -2 {object}  serializer.PureErrorResponse {"code":-2,"msg":"不合法的token"}
// @Failure -3 {object}  serializer.PureErrorResponse {"code":-3,"msg":"token过期"}
// @Router /ping [get]
func TokenCheck(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	j := middleware.JWT{}
	claims, err := j.ParseToken(token)
	if claims == nil || err != nil {
		c.JSON(http.StatusOK, serializer.PureErrorResponse{
			Status: -2,
			Msg:    "不合法的token",
		})
	} else if err == errors.New("Token is expired") {
		c.JSON(http.StatusOK, serializer.PureErrorResponse{
			Status: -3,
			Msg:    "Token过期",
		})
	} else if err != nil && err != errors.New("Token is expired") {
		c.JSON(http.StatusOK, serializer.PureErrorResponse{
			Status: -1,
			Msg:    err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, serializer.PureErrorResponse{
			Status: 0,
			Msg:    "success",
		})
	}

}
