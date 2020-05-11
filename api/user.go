package api

import (
	"github.com/gin-gonic/gin"
	"makespace-remaster/middleware"
	"makespace-remaster/service"
	"net/http"
)

// @Summary 用户注册
// @Description 用户注册
// @accept json
// @Produce  json
// @Param user body service.User true "user"
// @Success 200 {object} serializer.Response "{"status": 200,"data":null,"msg": "success" }"
// @Failure 4001 {object}  serializer.Response {"code":4001,"data":null,"msg":"JSON类型不匹配"}
// @Failure 4002 {object}  serializer.Response {"code":4002,"data":null,"msg":"参数错误"}
// @Failure 4003 {object}  serializer.Response {"code":4003,"data":null,"msg":"Tag Error"}
// @Router /registe [post]    //路由信息，一定要写上
func UserRegiste(c *gin.Context) {
	var service service.User
	if err := c.ShouldBindJSON(&service); err == nil {
		if response, err := service.Registe(); response != nil {
			c.JSON(http.StatusOK, response)
		} else {
			c.JSON(http.StatusOK, err)
		}
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}

}

// @Summary 用户登陆
// @Description 用户登陆
// @accept json
// @Produce  json
// @Param user body service.UserLoginService true "user"
// @Success 200 {object} 	serializer.Response "{"status": 200,"data":{"_id":000000,"token":token,"name":adeljck},"msg": "success" }"
// @Failure 4001 {object}  serializer.Response {"code":4001,"data":null,"msg":"JSON类型不匹配"}
// @Failure 4002 {object}  serializer.Response {"code":4002,"data":null,"msg":"参数错误"}
// @Failure 4003 {object}  serializer.Response {"code":4003,"data":null,"msg":"Tag Error"}
// @Failure 4005 {object}  serializer.Response {"code":4003,"data":null,"msg":"账号或密码错误"}
// @Failure 4006 {object}  serializer.Response {"code":4003,"data":null,"msg":"unauthrized account"}
// @Router /login [post]    //路由信息，一定要写上
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

// @Summary 企业入驻
// @Description 企业入驻
// @accept json
// @Produce  json
// @Param user body service.Enterprise true "enterprise"
// @Success 200 {object} 	serializer.Response "{"status":200, "msg": "success" }"
// @Failure 4001 {object}  serializer.Response {"code":4001,"data":null,"msg":"JSON类型不匹配"}
// @Failure 4002 {object}  serializer.Response {"code":4002,"data":null,"msg":"参数错误"}
// @Failure 4003 {object}  serializer.Response {"code":4003,"data":null,"msg":"Tag Error"}
// @Router /enterpriseapply [post]    //路由信息，一定要写上
func EnterpriseApply(c *gin.Context) {
	var service service.Enterprise
	if err := c.ShouldBindJSON(&service); err == nil {
		if success, err := service.Apply(); err != nil {
			c.JSON(http.StatusOK, success)
		} else {
			c.JSON(http.StatusOK, err)
		}
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func ForgePassword(c *gin.Context){

}
