package main

import (
	"makespace-remaster/conf"
	"makespace-remaster/server"
)

func main() {
	conf.Init()
	r := server.NewRouter()
	r.Run(":3000")
}
