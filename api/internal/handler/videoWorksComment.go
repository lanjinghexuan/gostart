package handler

import (
	"api/internal/server"
	pb "api/proto/videoWorksComment"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetCommentReq struct {
	WorksId  int `json:"worksId" form:"worksId"`
	Page     int `json:"page" form:"page"`
	Limit    int `json:"limit" form:"limit"`
	CommonId int `json:"commonId" form:"commonId"`
}

func GetComment(c *gin.Context) {
	var req GetCommentReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Limit == 0 {
		req.Limit = 10
	}
	if req.Page == 0 {
		req.Page = 1
	}

	res, err := server.GetComment(c, &pb.GetCommentReq{
		WorkId: int32(req.WorksId),
		Page:   int32(req.Page),
		Limit:  int32(req.Limit),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
	return
}

type GetCommentReplyListReq struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
	Pid   int `json:"pid" form:"pid" binding:"required"`
}

func GetCommentReplyList(c *gin.Context) {
	var req GetCommentReplyListReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Limit == 0 {
		req.Limit = 10
	}
	if req.Page == 0 {
		req.Page = 1
	}

	res, err := server.GetCommentReplyList(c, &pb.GetCommentReplyListReq{
		Pid:   int32(req.Pid),
		Page:  int32(req.Page),
		Limit: int32(req.Limit),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
	return
}
