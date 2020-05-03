package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"makespace-remaster/api"
	"makespace-remaster/middleware"
	"time"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.NoCache)
	r.Use(middleware.Options)
	r.Use(middleware.Secure)
	r.Use(middleware.RequestId())
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.GET("/api/v1/ping",api.Ping)
	v1 := r.Group("api/v1/")
	{
		v1.POST("/registe", api.UserRegiste)
		v1.POST("/login", api.UserLogin)
		v1.POST("/companyapply",api.UserLogin)
		authed := v1.Group("/", middleware.JWTAuth())
		{
			authed.POST("/logout", api.UserLogout)
		}
	}
	v1.GET("/project/list/*page",api.ProjectList)
	v1.POST("/check",api.TokenCheck)
	return r
}
