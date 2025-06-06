package handler

import (
	"api/internal/request"
	pb "api/proto/user"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"time"
)

func WorkAdd(c *gin.Context) {
	var req request.WorkAdd
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "验证失败",
			"data": err.Error(),
		})
		return
	}
	conn, err := grpc.NewClient("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	s := pb.NewUserClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	add, err := s.VideoWorkAdd(ctx, &pb.VideoWorkAddReq{
		Title:          req.Title,
		Describe:       req.Describe,
		MusicId:        int64(req.MusicId),
		Status:         req.Status,
		WorkCate:       req.WorkCate,
		Reviewer:       req.Reviewer,
		WorkPermission: req.WorkPermission,
		Address:        req.Address,
		LikeNum:        int64(req.LikeNum),
		ContentNum:     int64(req.ContentNum),
		ShareNum:       int64(req.ShareNum),
		CollectNum:     int64(req.CollectNum),
		BrowseNum:      int64(req.BrowseNum),
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "作品发布失败",
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "作品发布成功",
		"data": add,
	})
}

func WorkStatus(c *gin.Context) {
	var req request.WorkStatus
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "验证失败",
			"data": err.Error(),
		})
		return
	}
	conn, err := grpc.NewClient("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	s := pb.NewUserClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	status, err := s.WorkStatus(ctx, &pb.WorkStatusReq{
		WorkCate: req.WorkCate,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "作品发布失败",
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "作品发布成功",
		"data": status,
	})
}
