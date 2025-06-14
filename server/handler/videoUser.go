package handler

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"server/global"
	"server/inits/model"
	pb "server/proto/videoUser"
	"strconv"
	"time"
)

type VideoUserServer struct {
	pb.UnimplementedVideoUserServer
}

// 登录    获取验证码   验证验证码是否正确   查询用户信息是否存在(不存在需要注册)  返回用户id
func (s *VideoUserServer) Login(_ context.Context, in *pb.LoginReq) (*pb.LoginRes, error) {
	strMobile := strconv.Itoa(int(in.Mobile))
	code, err := global.REDIS.Get(global.Ctx, "code:"+strMobile).Result()
	if err != nil {
		fmt.Println("redis 获取不到验证码", err)
		return nil, err
	}
	if code != strconv.Itoa(int(in.Code)) {
		fmt.Println("redis 验证码错误", code)
		return nil, errors.New("验证码错误")
	}
	timeSend, err := global.REDIS.Get(global.Ctx, "send:time:"+strMobile).Result()
	if err != nil {
		fmt.Println("验证码过期", err)
		return nil, err
	}
	SendTime, err := strconv.Atoi(timeSend)

	if time.Now().Unix()-int64(SendTime) > 300 {
		return nil, errors.New("验证码已经过期")
	}
	var users model.VideoUser
	err = global.DB.Table("user").Where("mobile = ?", in.Mobile).Limit(1).Find(&users).Error
	if err != nil {
		fmt.Println("登录--查询用户信息失败", err)
		return nil, err
	}
	if users.Id == 0 {
		users = model.VideoUser{
			Mobile: strMobile,
			Name:   "默认名称",
		}
		err = global.DB.Table("user").Create(&users).Error
		if err != nil {
			fmt.Println("登录--添加用户信息失败", err)
			return nil, err
		}
	}

	return &pb.LoginRes{
		Id: users.Id,
	}, nil
}

// 获取用户信息     获取用户信息  查询用户头像 （接口会返回用户id ，跟被查看用户id。判断是否为自己的主页）
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
		UserId:        in.UserId,
	}, nil
}

// 点赞用户(多于功能)
func (s *VideoUserServer) Like(_ context.Context, in *pb.LikeReq) (*pb.LikeRes, error) {
	var like model.VideoLike
	err := global.DB.Table("video_like").Where("user_id = ?", in.UserId).Where("like_id = ?", in.LikeId).Limit(1).Find(&like).Error
	if err != nil {
		fmt.Println("查询信息出错：", err)
		return nil, status.Errorf(codes.Unknown, "查询信息出错")
	}
	if like.Id == 0 {
		like = model.VideoLike{
			UserId: in.UserId,
			LikeId: in.LikeId,
		}
		err = global.DB.Table("video_like").Create(&like).Error
		if err != nil {
			fmt.Println("点赞信息添加失败：", err)
			return nil, status.Errorf(codes.Unknown, "点赞信息添加失败")
		}
		return &pb.LikeRes{Success: true}, nil
	}
	err = global.DB.Table("user_like").Where("user_id = ?", in.UserId).Where("like_id = ?", in.LikeId).Limit(1).Delete(&like).Error
	if err != nil {
		fmt.Println("取消点赞失败：", err)
		return nil, status.Errorf(codes.Unknown, "取消点赞失败")
	}
	return &pb.LikeRes{Success: false}, nil
}

// 点赞视频/取消点赞视频
func (s *VideoUserServer) LikeVideo(_ context.Context, in *pb.LikeVideoReq) (*pb.LikeVideoRes, error) {
	var like model.VideoUserLike
	//查询是否被点赞过
	err := global.DB.Table("video_user_like").Where("user_id = ?", in.UserId).Where("video_id = ?", in.VideoId).Limit(1).Find(&like).Error
	if err != nil {
		fmt.Println("查询信息出错：", err)
		return nil, status.Errorf(codes.Unknown, "查询信息出错")
	}
	//没有点赞过 添加点赞
	if like.Id == 0 {
		like = model.VideoUserLike{
			UserId:  in.UserId,
			VideoId: in.VideoId,
		}
		err = global.DB.Transaction(func(tx *gorm.DB) error {
			err = tx.Table("video_user_like").Create(&like).Error
			if err != nil {
				return err
			}
			err = tx.Model("video_works").UpdateColumn("like_count ", gorm.Expr("like_count  + ?", 1)).Error
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return nil, errors.New("点赞失败")
		}
		return &pb.LikeVideoRes{Success: true}, nil
	}
	//点赞过取消点赞
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Table("video_user_like").Where("user_id = ?", in.UserId).Where("video_id = ?", in.VideoId).Limit(1).Delete(&like).Error
		if err != nil {
			return err
		}
		err = tx.Model("video_works").UpdateColumn("like_count ", gorm.Expr("like_count  - ?", 1)).Error
		if err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
	if err != nil {
		return nil, errors.New("取消点赞失败")
	}

	return &pb.LikeVideoRes{Success: false}, nil
}
