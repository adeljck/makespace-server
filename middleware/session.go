package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

// Session 初始化session
func Session() gin.HandlerFunc {
	//store session with redis
	store, _ := redis.NewStore(10000, "tcp", "www.adeljck.cn:6379", "weiruyi", []byte(os.Getenv("SESSIONSECRET")))
	redis.SetKeyPrefix(store, "hnit_")
	//store := cookie.NewStore([]byte(secret))
	//Also set Secure: true if using SSL, you should though
	store.Options(sessions.Options{HttpOnly: true, MaxAge: int(30 * time.Minute), Path: "/", Domain: "127.0.0.1"})
	return sessions.Sessions(os.Getenv("SESSIONNAME"), store)
}
