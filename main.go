package main

import (
	"ginStudy/internal/initialize"
)

func main() {
	server := initialize.InitServer()
	err := server.Run(":80")
	if err != nil {
		panic("服务启动失败")
	}
}
