package initialize

import (
	"ginStudy/internal/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitServer() *gin.Engine {
	server := gin.Default()

	//解决跨域
	//server.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"http://127.0.0.1"},
	//	AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
	//	AllowHeaders:     []string{"Origin"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//	AllowOriginFunc: func(origin string) bool {
	//		return origin == "https://github.com"
	//	},
	//	MaxAge: 12 * time.Hour,
	//}))
	server.Use(cors.Default())

	baseRouter := server.Group("/api/v1")
	router.InitUserRouter(baseRouter)

	return server
}
