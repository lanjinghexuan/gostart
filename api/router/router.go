package router

import "github.com/gin-gonic/gin"

func LoadRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		UserRouter(api)
		WorkRouter(api)
	}
}
