package router

import (
	"api/internal/handler"
	"api/middlewear"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	c := r.Group("/api")
	{
		c.POST("/Login", handler.Login)                            // 登录
		c.GET("/SendCode", handler.SendCode)                       // 获取验证码
		c.GET("/getVideoWorks", handler.GetVideoWorks)             //获取视频列表
		c.GET("/getComment", handler.GetComment)                   //获取视频评论
		c.GET("/getCommentReplyList", handler.GetCommentReplyList) // 获取子评论

		c.Use(middlewear.Jwt())                         // 校验用户是否登录(token)
		c.GET("/getUserInfo", handler.GetUserInfo)      // 获取用户信息
		c.GET("/Like", handler.Like)                    // 点赞用户(多余)
		c.GET("LikeVideo", handler.LikeVideo)           // 点赞视频 (第二次取消)
		c.POST("/addVideoWorks", handler.AddVideoWorks) // 添加作品
	}
}
