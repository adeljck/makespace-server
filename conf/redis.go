package conf

import (
	"github.com/gomodule/redigo/redis"
	"os"
	"strconv"
	"time"
)

var (
	RedisPool   *redis.Pool
	maxIdle     int
	maxActive   int
	idleTimeout int
	timeout     int
)

func getvalue() {
	maxIdle, _ = strconv.Atoi(os.Getenv("MAX_IDLE"))
	maxActive, _ = strconv.Atoi(os.Getenv("MAX_ACTIVE"))
	idleTimeout, _ = strconv.Atoi(os.Getenv("IDLE_TIMEOUT"))
	timeout, _ = strconv.Atoi(os.Getenv("TIME_OUT"))
}
func RedisInit() {

	RedisPool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: time.Duration(idleTimeout) * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", os.Getenv("REDIS_HOST"),
				redis.DialPassword(os.Getenv("REDIS_PASSWORD")),
				redis.DialConnectTimeout(time.Duration(timeout)*time.Second),
				redis.DialReadTimeout(time.Duration(timeout)*time.Second),
				redis.DialWriteTimeout(time.Duration(timeout)*time.Second))
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}
}
