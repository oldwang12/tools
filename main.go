package main

import (
	"github.com/oldwang12/tools/pkg/routes"
)

func main() {
	r := routes.InitRouter()
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
