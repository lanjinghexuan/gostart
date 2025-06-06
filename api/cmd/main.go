package main

import (
	"api/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.LoadRouter(r)
	r.Run(":8081") // 监听并在 0.0.0.0:8080 上启动服务
}
2