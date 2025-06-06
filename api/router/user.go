package router

import (
	"api/internal/handler"
	"github.com/gin-gonic/gin"
)

func UserRouter(api *gin.RouterGroup) {
	user := api.Group("/user")
	{
		user.POST("/sendSms", handler.SendSms)
		user.POST("/register", handler.Register)
		user.POST("/login", handler.Login)
		user.POST("/UserCenterAdd", handler.UserCenterAdd)
	}
}
