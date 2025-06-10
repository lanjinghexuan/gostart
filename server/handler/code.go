package handler

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"server/global"
	pb "server/proto/code"
	"time"
)

type CodeServer struct {
	pb.UnimplementedCodeServer
}

func (c *CodeServer) SendCode(ctx context.Context, req *pb.SendCodeReq) (*pb.SendCodeRes, error) {
	t, err := global.REDIS.Get(global.Ctx, "time"+req.Mobile).Result()
	if err != redis.Nil || t != "" {
		return nil, errors.New("60秒只能发送一次验证码")
	}
	err = global.REDIS.Set(global.Ctx, "code:"+req.Mobile, req.Code, 1*time.Hour).Err()
	if err != nil {
		return nil, err
	}
	global.REDIS.Set(global.Ctx, "time:"+req.Mobile, req.Code, 60*time.Second)
	return &pb.SendCodeRes{Success: true}, nil
}
