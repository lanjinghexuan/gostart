package handler

import (
	"context"
	"gorm.io/gorm"
	"server/global"
	"server/inits/model"
	pb "server/proto/videoWorks"
	"time"
)

type VideoWorksServer struct {
	pb.UnimplementedVideoWorksServer
}

// 获取视频列表   分页
func (c *VideoWorksServer) GetVideoWorks(ctx context.Context, req *pb.GetVideoWorksReq) (res *pb.GetVideoWorksRes, err error) {

	limit := 10
	start := (int(req.Page) - 1) * limit
	var works []*model.VideoWorks
	//获取视频
	err = global.DB.Table("video_works").Limit(limit).Offset(start).Find(&works).Error
	if err != nil {
		return &pb.GetVideoWorksRes{}, err
	}
	//获取关联的音乐id
	var music_id []int16
	var work_id []int32
	for _, v := range works {
		music_id = append(music_id, v.MusicId)
		work_id = append(work_id, v.Id)
	}
	//获取关联的视频地址id
	var wordObject []*model.VideoWorkObject
	global.DB.Table("video_work_object").Where("work_id in ?", work_id).Find(&wordObject)
	var object_id []int16
	for _, v := range wordObject {
		object_id = append(object_id, v.ObjectId)
	}
	//获取关联的音乐跟视频地址
	var musics_all []*model.VideoMusic
	global.DB.Table("video_music").Where("id in (?)", music_id).Find(&musics_all)
	var object_all []*model.VideoObject
	global.DB.Table("video_object").Where("id in (?)", object_id).Find(&object_all)

	//格式转换为map方便匹配
	musicMap := make(map[int32]*model.VideoMusic)
	for _, v := range musics_all {
		musicMap[v.Id] = v
	}
	objectMap := make(map[int32]*model.VideoObject)
	for _, v := range object_all {
		objectMap[v.Id] = v
	}
	objectMaps := make(map[int32]*model.VideoObject)
	for _, v := range wordObject {
		objectMaps[int32(v.WorkId)] = objectMap[int32(v.ObjectId)]
	}

	var list []*pb.VideoWorksData
	//循环处理需要返回的数据格式(头像跟music可能因为脏数据给了默认值)
	for _, v := range works {
		var music *model.VideoMusic
		if mus, ok := musicMap[int32(v.MusicId)]; ok {
			music = mus
		} else {
			music = &model.VideoMusic{
				Id:   int32(1),
				Name: "默认音乐",
				Path: "https://1258321826.oss-cn-shanghai.aliyuncs.com/%E9%81%BF%E8%84%B8.mp3",
			}
		}

		var object *model.VideoObject
		if mus, ok := objectMaps[v.Id]; ok {
			object = mus
		} else {
			object = &model.VideoObject{
				Id:   0,
				Path: "",
			}
		}

		list = append(list, &pb.VideoWorksData{
			Id:           v.Id,
			Title:        v.Title,
			Desc:         v.Desc,
			Music:        music.Path,
			MusicName:    music.Name,
			ObjectPath:   object.Path,
			LikeCount:    int32(v.LikeCount),
			CommentCount: int32(v.CommentCount),
		})
	}

	return &pb.GetVideoWorksRes{
		Data: list,
	}, nil
}

func (c *VideoWorksServer) AddVideoWorks(ctx context.Context, req *pb.AddVideoWorksReq) (res *pb.AddVideoWorksRes, err error) {
	data := &model.VideoWorks{
		Title:          req.Title,
		Desc:           req.Desc,
		MusicId:        int16(req.MusicId),
		WorkType:       req.WorkType,
		CheckStatus:    "1",
		CheckUser:      int16(req.UserId),
		IpAddress:      req.Ip,
		WorkPermission: "",
		LikeCount:      0,
		CommentCount:   0,
		ShareCount:     0,
		CollectCount:   0,
		BrowseCount:    0,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
	}
	addObject := &model.VideoObject{
		Path: req.FilePath,
	}

	err = global.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Table("video_object").Create(&addObject).Error
		if err != nil {
			return err
		}

		err = tx.Table("video_works").Create(&data).Error
		if err != nil {
			return err
		}

		arr := model.VideoWorkObject{
			WorkId:    int16(data.Id),
			ObjectId:  int16(addObject.Id),
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}

		err = tx.Table("video_work_object").Create(&arr).Error
		if err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})
	if err != nil {
		return &pb.AddVideoWorksRes{}, err
	}
	return &pb.AddVideoWorksRes{
		Succ: true,
	}, nil
}
