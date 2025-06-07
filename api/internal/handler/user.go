package handler

import (
	"api/internal/request"
	"api/pkg"
	pb "api/proto/user"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"time"
)

func SendSms(c *gin.Context) {
	var req request.SendSms
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
	sendSms, err := s.SendSms(ctx, &pb.SendSmsReq{
		Phone:  req.Phone,
		Source: req.Source,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "注册失败",
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "短信发送成功",
		"data": sendSms,
	})
}
func Register(c *gin.Context) {
	var req request.Register
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
	register, err := s.Register(ctx, &pb.RegisterReq{
		Phone:    req.Phone,
		Password: req.Password,
		SmsCode:  req.SmsCode,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "注册失败",
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
		"data": register,
	})
}
func Login(c *gin.Context) {
	var req request.Login
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
	login, err := s.Login(ctx, &pb.LoginReq{
		Phone:    req.Phone,
		Password: req.Password,
		SmsCode:  req.SmsCode,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "登录失败",
			"data": err.Error(),
		})
		return
	}
	token, err := pkg.NewJWT("2210a").CreateToken(pkg.CustomClaims{
		ID: uint(login.Id),
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "生成token失败",
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": token,
	})
}
func UserCenterAdd(c *gin.Context) {
	var req request.UserCenterAdd
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
	add, err := s.UserCenterAdd(ctx, &pb.UserCenterAddReq{
		Name:          req.Name,
		NickName:      req.NickName,
		UserCode:      req.UserCode,
		Signature:     "",
		Sex:           "",
		IpAddress:     "",
		Constellation: "",
		AttendCount:   0,
		FanCount:      0,
		ZanCount:      0,
		Status:        "",
		ImageId:       0,
		Info:          "",
		Mobile:        "",
		RealNameAuth:  "",
		Age:           "",
		OnlineStatus:  "",
		Type:          "",
		Level:         "",
		Balance:       "",
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "注册失败",
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
		"data": add,
	})
}
