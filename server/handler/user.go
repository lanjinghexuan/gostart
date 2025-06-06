package handler

import (
	"context"
	"math/rand"
	"server/global"
	"server/inits/model"
	"server/pkg"
	pb "server/proto/user"
	"strconv"
	"time"
)

type UserServer struct {
	pb.UnimplementedUserServer
}

func (s *UserServer) SendSms(_ context.Context, in *pb.SendSmsReq) (*pb.SendSmsResp, error) {
	code := rand.Intn(9000) + 1000
	sms, err := pkg.SendSms(in.Phone, strconv.Itoa(code))
	if err != nil {
		return nil, err
	}
	if *sms.Body.Code != "ok" {
		return &pb.SendSmsResp{
			Code: 500,
			Msg:  "失败",
		}, nil
	}
	global.RedisDb.Set(context.Background(), "sendSms"+in.Source+in.Phone, code, time.Minute*5)
	return &pb.SendSmsResp{
		Code: 200,
		Msg:  "短信发送成功",
	}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *UserServer) Register(_ context.Context, in *pb.RegisterReq) (*pb.RegisterResp, error) {

	user := model.User{
		Phone:    in.Phone,
		Password: in.Password,
	}
	err := global.DB.Debug().Create(&user).Error
	if err != nil {
		return &pb.RegisterResp{
			Code: 500,
			Msg:  "注册失败",
		}, nil
	}
	get := global.RedisDb.Get(context.Background(), "sendSms"+"register"+in.Phone)
	if in.SmsCode != get.Val() {
		return &pb.RegisterResp{
			Code: 500,
			Msg:  "验证码错误",
		}, nil
	}
	return &pb.RegisterResp{
		Code: 200,
		Msg:  "注册成功",
	}, nil
}
func (s *UserServer) Login(_ context.Context, in *pb.LoginReq) (*pb.LoginResp, error) {
	var user model.User
	err := global.DB.Debug().Where("phone = ?", in.Phone).Find(&user).Error
	if err != nil {
		return &pb.LoginResp{
			Code: 500,
			Msg:  "查询失败",
		}, nil
	}
	if in.Phone == "" {
		return &pb.LoginResp{
			Code: 500,
			Msg:  "用户不存在",
		}, nil
	}
	if in.Password != user.Password {
		return &pb.LoginResp{
			Code: 500,
			Msg:  "密码错误",
		}, nil
	}
	get := global.RedisDb.Get(context.Background(), "sendSms"+"login"+in.Phone)
	if in.SmsCode != get.Val() {
		return &pb.LoginResp{
			Code: 500,
			Msg:  "验证码错误",
		}, nil
	}
	return &pb.LoginResp{
		Code: 200,
		Msg:  "登录成功",
		Id:   int64(user.ID),
	}, nil
}

func (s *UserServer) VideoWorkAdd(_ context.Context, in *pb.VideoWorkAddReq) (*pb.VideoWorkAddResp, error) {
	work := model.Work{
		Title:          in.Title,
		Describe:       in.Describe,
		MusicId:        int(in.MusicId),
		Status:         in.Status,
		WorkCate:       in.WorkCate,
		Reviewer:       in.Reviewer,
		WorkPermission: in.WorkPermission,
		Address:        in.Address,
		LikeNum:        int(in.LikeNum),
		ContentNum:     int(in.ContentNum),
		ShareNum:       int(in.ShareNum),
		CollectNum:     int(in.CollectNum),
		BrowseNum:      int(in.BrowseNum),
	}
	err := global.DB.Debug().Create(&work).Error
	if err != nil {
		return &pb.VideoWorkAddResp{
			Code: 500,
			Msg:  "作品发布失败",
		}, nil
	}
	return &pb.VideoWorkAddResp{
		Code: 200,
		Msg:  "作品发布成功",
	}, nil
}

func (s *UserServer) WorkStatus(_ context.Context, in *pb.WorkStatusReq) (*pb.WorkStatusResp, error) {
	var works []model.Work
	result := global.DB.Debug().Where("work_cate = ?", in.WorkCate).Find(&works)
	if result.Error != nil {
		return nil, result.Error
	}
	var lists []*pb.WorkStatus
	for _, m := range works {
		list := &pb.WorkStatus{
			Title:          m.Title,
			Describe:       m.Describe,
			MusicId:        int64(m.MusicId),
			Status:         m.Status,
			WorkCate:       m.WorkCate,
			Reviewer:       m.Reviewer,
			WorkPermission: m.WorkPermission,
			Address:        m.Address,
			LikeNum:        int64(m.LikeNum),
			ContentNum:     int64(m.ContentNum),
			ShareNum:       int64(m.ShareNum),
			CollectNum:     int64(m.CollectNum),
			BrowseNum:      int64(m.BrowseNum),
		}
		lists = append(lists, list)
	}
	return &pb.WorkStatusResp{
		List: lists,
	}, nil
}
