package amdin

import (
	"github.com/gin-gonic/gin"
	"makespace-remaster/api"
	"makespace-remaster/middleware"
	"makespace-remaster/service/amdin"
	"net/http"
)

func AdminLogin(c *gin.Context) {
	var service amdin.AdminLoginService
	if err := c.ShouldBindJSON(&service); err == nil {
		if info, err := service.Login(); err != nil {
			c.JSON(http.StatusOK, err)
		} else {
			middleware.GenerateToken(c, info)
		}
	} else {
		c.JSON(http.StatusOK, api.ErrorResponse(err))
	}
}
