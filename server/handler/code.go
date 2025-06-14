package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"server/global"
	pb "server/proto/code"
	"time"
)

type CodeServer struct {
	pb.UnimplementedCodeServer
}

// 发送验证码    验证手机号格式在api处理     限制60S发送
func (c *CodeServer) SendCode(ctx context.Context, in *pb.SendCodeReq) (*pb.SendCodeRes, error) {
	sendTime, err := global.REDIS.Get(global.Ctx, "time:"+in.Mobile).Result()
	if err != redis.Nil {
		fmt.Println(sendTime, err)
		if err != nil {
			fmt.Println(sendTime, err)
			return nil, err
		}

	}
	fmt.Println(sendTime)
	if sendTime != "" {
		return nil, errors.New("60秒发送一次")
	}
	code := rand.Intn(9000) + 1000
	global.REDIS.Set(global.Ctx, "code:"+in.Mobile, code, 0)
	global.REDIS.Set(global.Ctx, "send:time:"+in.Mobile, time.Now().Unix(), 0)
	global.REDIS.Set(global.Ctx, "time:"+in.Mobile, code, time.Minute*1)
	fmt.Println("code:", code, "time:", time.Now().Unix())
	return &pb.SendCodeRes{Success: true}, nil
}
