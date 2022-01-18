package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//创建一个router实例
	router := gin.Default()

	//路由
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//默认监听并在 0.0.0.0:8080 上启动服务
	router.Run()
}
