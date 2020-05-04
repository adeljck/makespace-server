package main

import (
	"makespace-remaster/conf"
	"makespace-remaster/server"
)

// @title 湖南工学院创客系统API
// @version 1.0
// @description 湖南工学院创客系统API介绍
// @termsOfService http://www.hnit.edu.cn

// @contact.name adeljck
// @contact.url http://adeljck.github.io
// @contact.email fuck@you.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:3000
// @BasePath /api/v1
func main() {
	conf.Init()
	r := server.NewRouter()
	r.Run(":3000")
}
