package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	server := gin.Default()
	//解决跨域
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	baseRouter := server.Group("/api/v1")
	ug := baseRouter.Group("/user")
	{
		ug.POST("/login", login)
	}

	server.Run(":80")
}

func login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "请求成功",
		"data": "xxxx",
	})
}
