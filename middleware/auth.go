package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"makespace-remaster/serializer"
	"net/http"
)

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionValue := session.Get("username")
		if sessionValue == nil {
			c.JSON(http.StatusUnauthorized, serializer.Response{
				Status: 4001,
				Msg:    "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Set("username", sessionValue.(string))
		c.Next()
		return
	}
}

func SaveAuthSession(c *gin.Context, username string) {
	session := sessions.Default(c)
	session.Clear()
	session.Set("username", username)
	session.Save()
}
func ClearAuthSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}
func HasSession(c *gin.Context) bool {
	session := sessions.Default(c)
	if sessionValue := session.Get("username"); sessionValue == nil {
		return false
	}
	return true
}
func GetSessionUserId(c *gin.Context) string {
	session := sessions.Default(c)
	sessionValue := session.Get("username")
	if sessionValue == nil {
		return ""
	}
	return sessionValue.(string)

}
