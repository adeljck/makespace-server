package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"makespace-remaster/conf"
	"makespace-remaster/middleware"
	"makespace-remaster/serializer"
	"makespace-remaster/service"
	"net/http"
)

func UserRegiste(c *gin.Context) {
	var service service.User
	if err := c.ShouldBindJSON(&service); err == nil {
		if response,err := service.Registe(); response != nil {
			c.JSON(http.StatusOK,response)
		}else{
			c.JSON(http.StatusOK,err)
		}
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}

}

func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBindJSON(&service); err == nil {
		if info, err := service.Login(); err != nil {
			c.JSON(http.StatusOK, err)
		} else {
			middleware.GenerateToken(c, info)
		}
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	j := middleware.NewJWT()
	claims, _ := j.ParseToken(c.Request.Header.Get("Athorization"))
	conn, _ := conf.RedisPool.Dial()
	conn.Do("DEL", claims.Id)
	c.JSON(200, serializer.PureErrorResponse{
		Status: 0,
		Msg:    "登出成功",
	})
}
func TokenCheck(c *gin.Context){
	j := middleware.NewJWT()
	claims, _ := j.ParseToken(c.Request.Header.Get("Athorization"))
	fmt.Println(claims)
}