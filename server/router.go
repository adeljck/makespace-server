package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"makespace-remaster/api"
	"makespace-remaster/api/amdin"
	_ "makespace-remaster/docs"
	"makespace-remaster/middleware"
	"os"
	"time"
)

func NewRouter() *gin.Engine {
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.Default()
	r.Use(middleware.Cors())
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
	v1 := r.Group("api/v1/")
	{
		v1.GET("/academy", api.GetAcademy)
		v1.GET("/industry", api.GetIndustry)
		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		v1.GET("/ping", api.Ping)
		v1.POST("/registe", api.UserRegiste)
		v1.POST("/enregiste", api.EnterpriseUserRegiste)
		v1.POST("/login", api.UserLogin)
		v1.POST("/enterpriseapply", api.EnterpriseApply)
		v1.POST("/forget", api.ForgePassword)
		v1.POST("/enforget", api.EnterpriseForgePassword)
		v1.GET("/project", api.GetProjectInfo)
		v1.GET("/token/check", api.TokenCheck)
		v1.GET("/projects/*page", api.ProjectList)
		v1.GET("/activities/*page", api.ActivityList)
		v1.GET("/activity/", api.GetActivityInfo)
		v1.GET("/projecttype", api.GetProjectType)
		v1.GET("/captcha", func(c *gin.Context) {
			api.Captcha(c, 4)
		})
		v1.POST("/captcha", api.CaptchaVerify)
		admin := v1.Group("/admin", amdin.AdminLogin)
		{
			admin.POST("/login")
		}
		authed := v1.Group("/", middleware.JWTAuth())
		{
			authed.POST("/forget/changepassword", api.ForgetChangePassword)
			authed.POST("/enforget/changepassword", api.EnterpriseForgetChangePassword)
			authed.POST("/mine/changepassword", api.ChangePassword)
			authed.DELETE("/project", api.DeleteProject)
			authed.PUT("/project", api.GetProject)
			authed.POST("/project", api.AddProject)
			authed.DELETE("/activity", api.DeleteActivity)
			authed.PUT("/activity", api.GetActivity)
			authed.POST("/activity", api.AddActivity)
			//authed.POST("/academy", api.AddAcademy)
			//authed.DELETE("/academy", api.DeleteAcademy)
			//authed.PUT("/academy", api.UpdateAcademy)
			//authed.POST("/industry", api.AddIndustry)
			//authed.DELETE("/industry", api.DeleteIndustryy)
			//authed.PUT("/industry", api.UpdateIndustry)
		}
	}
	return r
}
