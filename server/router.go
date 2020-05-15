package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"makespace-remaster/api"
	_ "makespace-remaster/docs"
	"makespace-remaster/middleware"
	"os"
	"time"
)

func NewRouter() *gin.Engine {
	gin.SetMode(os.Getenv("GIN_MODE"))
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
	r.GET("/api/v1/ping", api.Ping)
	r.GET("/api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("api/v1/")
	{
		v1.POST("/registe", api.UserRegiste)
		v1.POST("/login", api.UserLogin)
		v1.POST("/enterpriseapply", api.EnterpriseApply)
		//authed := v1.Group("/", middleware.JWTAuth())
		//{
		//	authed.POST("/check",api.TokenCheck)
		//}
	}
	v1.GET("/project/list/*page", api.ProjectList)
	v1.POST("/project/search", api.SearchProject)
	return r
}
