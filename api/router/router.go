package router

import (
	"api/internal/handler"
	"api/middlewear"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	c := r.Group("/api")
	{
		c.POST("/Login", handler.Login)
		c.GET("/SendCode", handler.SendCode)

		c.Use(middlewear.Jwt())
		c.GET("/getUserInfo", handler.GetUserInfo)
	}
}
