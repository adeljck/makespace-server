package api

import (
	"github.com/gin-gonic/gin"
	"makespace-remaster/serializer"
)

// @Summary 服务器状态检查
// @Description 服务器状态检查
// @accept */*
// @Produce  json
// @Success 200 {object} serializer.Response "{"status": 0,"data": null,"msg": "Pong" }"
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Status: 0,
		Msg:    "Pong",
		Data:   nil,
	})
}
