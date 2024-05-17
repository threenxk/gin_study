package router

import (
	"ginStudy/internal/web"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(router *gin.RouterGroup) {
	handler := web.NewUserHandler()
	ug := router.Group("/user")
	{
		ug.POST("/login", handler.Login)
	}
}
