package handler

import (
	"api/internal/server"
	"api/pkg"
	pb "api/proto/videoUser"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginReq struct {
	Mobile int64 `json:"mobile" form:"mobile" binding:"required"`
	Code   int32 `json:"code" form:"code" binding:"required"`
}

func Login(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := server.Login(c, &pb.LoginReq{
		Mobile: req.Mobile,
		Code:   req.Code,
	})
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if userId.Id == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "用户不存在"})
		return
	}

	token, err := pkg.GetToken(userId.Id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"code":  200,
	})
	return
}

//type GetUserInfoReq struct {
//	UserId int32 `json:"userId" form:"userId" binding:"required"`
//}

func GetUserInfo(c *gin.Context) {
	userId := c.MustGet("userId")
	fmt.Println(userId)
	//var req GetUserInfoReq
	//if err := c.ShouldBind(&req); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}

	data, err := server.GetUserInfo(c, &pb.GetUserInfoReq{
		//UserId: req.UserId,
		UserId: userId.(int32),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
		"code": 200,
	})
	return
}

type LikeReq struct {
	LikeId int32 `json:"like_id" form:"like_id"`
}

func Like(c *gin.Context) {
	var req LikeReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
	}

	res, err := server.Like(c, &pb.LikeReq{
		UserId: c.MustGet("userId").(int32),
		LikeId: req.LikeId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	if res == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 801,
			"msg":  "系统出错",
		})
		return
	}

	if res.Success {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "点赞成功",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "取消点赞成功",
	})
	return
}

type LikeVideoReq struct {
	VideoId int32 `json:"video_id" form:"video_id" binding:"required"`
}

// 用户点赞视频 重复点击会导致点赞取消
func LikeVideo(c *gin.Context) {
	//获取用户点赞的视频id
	var req LikeVideoReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	//请求服务端 userid从中间件获取
	res, err := server.LikeVideo(c, &pb.LikeVideoReq{
		UserId:  c.MustGet("userId").(int32),
		VideoId: req.VideoId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	if res == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 801,
			"msg":  "系统出错",
		})
		return
	}

	if res.Success {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "点赞成功",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "取消点赞成功",
	})
	return
}
