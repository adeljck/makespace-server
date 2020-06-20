package api

import (
	"github.com/gin-gonic/gin"
	"makespace-remaster/middleware"
	"makespace-remaster/serializer"
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
func EnterpriseUserRegiste(c *gin.Context) {
	var service service.EnterpriseUser
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

// @Summary 忘记密码
// @Description 忘记密码
// @accept json
// @Produce  json
// @Param user body service.UserForget true "user"
// @Success 200 {object} 	serializer.Response "{"status": 200,"data":{"_id":000000,"token":token,"name":adeljck},"msg": "success" }"
// @Failure 4001 {object}  serializer.Response {"code":4001,"data":null,"msg":"JSON类型不匹配"}
// @Failure 4002 {object}  serializer.Response {"code":4002,"data":null,"msg":"参数错误"}
// @Failure 4003 {object}  serializer.Response {"code":4003,"data":null,"msg":"Tag Error"}
// @Failure 4006 {object}  serializer.Response {"code":4003,"data":null,"msg":"unauthrized account"}
// @Router /forget [post]    //路由信息，一定要写上
func ForgePassword(c *gin.Context) {
	var service service.UserForget
	if err := c.ShouldBind(&service); err == nil {
		if info, err := service.Forget(); info != nil {
			middleware.GenerateToken(c, info)
		} else {
			c.JSON(http.StatusOK, err)
		}
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}
func EnterpriseForgePassword(c *gin.Context) {
	var service service.EnterpriseUserForget
	if err := c.ShouldBind(&service); err == nil {
		if info, err := service.Forget(); info != nil {
			middleware.GenerateToken(c, info)
		} else {
			c.JSON(http.StatusOK, err)
		}
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}
func EnterpriseForgetChangePassword(c *gin.Context) {
	var service service.EnterpriseForgetChangePassword
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(200, serializer.PureErrorResponse{
			Status: 5001,
			Msg:    "未登陆",
		})
	}
	j := middleware.JWT{}
	claims, _ := j.ParseToken(token)
	if err := c.ShouldBind(&service); err == nil {
		if res, err := service.Change(claims.Id); res != nil {
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusOK, err)
		}
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}
// @Summary 忘记密码->修改密码
// @Description 忘记密码->修改密码
// @accept json
// @Produce  json
// @Param user body service.UserForgetChangePassword true "user"
// @Success 200 {object} 	serializer.Response "{"status": 200,"msg": "success" }"
// @Failure 4001 {object}  serializer.Response {"code":4001,"data":null,"msg":"JSON类型不匹配"}
// @Failure 4002 {object}  serializer.Response {"code":4002,"data":null,"msg":"参数错误"}
// @Failure 4003 {object}  serializer.Response {"code":4003,"data":null,"msg":"Tag Error"}
// @Failure 4006 {object}  serializer.Response {"code":4003,"data":null,"msg":"unauthrized account"}
// @Router /forget/changepassword [post]    //路由信息，一定要写上
func ForgetChangePassword(c *gin.Context) {
	var service service.UserForgetChangePassword
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(200, serializer.PureErrorResponse{
			Status: 5001,
			Msg:    "未登陆",
		})
	}
	j := middleware.JWT{}
	claims, _ := j.ParseToken(token)
	if err := c.ShouldBind(&service); err == nil {
		if res, err := service.Change(claims.Id); res != nil {
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusOK, err)
		}
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// @Summary 用户修改密码
// @Description 用户修改密码
// @accept json
// @Produce  json
// @Param user body service.UserChangePassword true "user"
// @Success 200 {object} 	serializer.Response "{"status": 200,"msg": "success" }"
// @Failure 4001 {object}  serializer.Response {"code":4001,"data":null,"msg":"JSON类型不匹配"}
// @Failure 4002 {object}  serializer.Response {"code":4002,"data":null,"msg":"参数错误"}
// @Failure 4003 {object}  serializer.Response {"code":4003,"data":null,"msg":"Tag Error"}
// @Failure 4006 {object}  serializer.Response {"code":4003,"data":null,"msg":"unauthrized account"}
// @Router /mine/changepassword [post]    //路由信息，一定要写上
func ChangePassword(c *gin.Context) {
	var service service.UserChangePassword
	token := c.Request.Header.Get("Athorization")
	j := middleware.JWT{}
	claims, _ := j.ParseToken(token)
	if err := c.ShouldBind(&service); err == nil {
		if res, err := service.Change(claims.Id); res != nil {
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusOK, err)
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
	token := c.Request.Header.Get("Athorization")
	if token == "" {
		c.JSON(200, serializer.PureErrorResponse{
			Status: 5001,
			Msg:    "未登陆",
		})
	}
	j := middleware.JWT{}
	claims, _ := j.ParseToken(token)
	if err := c.ShouldBindJSON(&service); err == nil {
		if success, err := service.Apply(claims.Role); err == nil {
			c.JSON(http.StatusOK, success)
		} else {
			c.JSON(http.StatusOK, err)
		}
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}
