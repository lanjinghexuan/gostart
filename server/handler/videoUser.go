package handler

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"server/global"
	"server/inits/model"
	pb "server/proto/videoUser"
	"strconv"
	"time"
)

type VideoUserServer struct {
	pb.UnimplementedVideoUserServer
}

func (s *VideoUserServer) Login(_ context.Context, in *pb.LoginReq) (*pb.LoginRes, error) {

	ttlTime, err := global.REDIS.TTL(global.Ctx, "code:"+strconv.FormatInt(in.Mobile, 10)).Result()

	if err != nil || ttlTime < time.Minute*55 {
		return nil, status.Errorf(codes.Unknown, "验证码已经过期")
	}
	code, _ := global.REDIS.Get(global.Ctx, "code:"+strconv.FormatInt(in.Mobile, 10)).Result()
	if code != strconv.Itoa(int(in.Code)) {
		return nil, status.Errorf(codes.Unknown, "验证码不正确")
	}

	var models model.VideoUser
	err = global.DB.Table("video_user").Where("mobile = ?", in.Mobile).Limit(1).Find(&models).Error
	if err != nil {
		return nil, err
	}
	if models.Id == 0 {
		addModels := model.VideoUser{
			Mobile: strconv.Itoa(int(in.Mobile)),
		}
		global.DB.Table("video_user").Create(&addModels)
		models = addModels
		fmt.Println(models)
	}
	return &pb.LoginRes{Id: models.Id}, nil
}

func (s *VideoUserServer) GetUserInfo(_ context.Context, in *pb.GetUserInfoReq) (*pb.GetUserInfoRes, error) {
	var models model.VideoUser
	imgurl := "https://1258321826.oss-cn-shanghai.aliyuncs.com/b2.png"
	err := global.DB.Table("video_user").Where("id = ?", in.UserId).Limit(1).Find(&models).Error
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "查询失败")
	}
	if models.AvatorFileId != 0 {
		var avator model.VideoAvator
		err := global.DB.Table("video_avator").Where("id = ?", models.AvatorFileId).Limit(1).Find(&avator).Error
		if err != nil {
			return nil, status.Errorf(codes.Unknown, "查询失败")
		}
		if avator.Id != 0 {
			imgurl = avator.Avator
		}

	}

	return &pb.GetUserInfoRes{
		Avator:        imgurl,
		NickName:      models.NickName,
		Id:            models.Id,
		Signature:     models.Signature,
		Constellation: models.Constellation,
		IpAddress:     models.IpAddress,
	}, nil
}
