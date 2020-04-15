package conf

import (
	"github.com/joho/godotenv"
	"makespace-remaster/module"
)

func Init() {
	godotenv.Load()
	RedisInit()
	module.MongoInit()
	OssInit()
}
