package handler

import (
	"context"
	"server/global"
	"server/inits/model"
	pb "server/proto/videoUser"
)

type server struct {
	pb.UnimplementedVideoUserServer
}

func (s *server) Login(_ context.Context, in *pb.LoginReq) (*pb.LoginRes, error) {
	var models model.VideoUser
	err := global.DB.Table("video_user").Where("mobile = ?", in.Mobile).First(&models).Error
	if err != nil {
		return nil, err
	}
	return &pb.LoginRes{Id: models.Id}, nil
}
