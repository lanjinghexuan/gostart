package router

import (
	"api/internal/handler"
	"github.com/gin-gonic/gin"
)

func WorkRouter(api *gin.RouterGroup) {
	work := api.Group("/work")
	{
		work.POST("/WorkAdd", handler.WorkAdd)
		work.POST("/WorkStatus", handler.WorkStatus)
	}
}
