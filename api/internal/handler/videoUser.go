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
