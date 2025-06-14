package handler

import (
	"context"
	"fmt"
	"server/global"
	"server/inits/model"
	pb "server/proto/videoWorksComment"
	"strings"
)

type VideoWorksCommentServer struct {
	pb.UnimplementedVideoWorksCommentServer
}

func (c VideoWorksCommentServer) GetComment(ctx context.Context, in *pb.GetCommentReq) (*pb.GetCommentRes, error) {
	start := (in.Page - 1) * in.Limit

	// 1. 获取主评论（pid=0）
	var mainComments []*model.VideoWorkComment
	err := global.DB.Table("video_work_comment").
		Where("work_id = ? AND pid = 0", in.WorkId).
		Offset(int(start)).Limit(int(in.Limit)).
		Order("created_at ASC").
		Find(&mainComments).Error
	if err != nil {
		return &pb.GetCommentRes{}, err
	}

	if len(mainComments) == 0 {
		return &pb.GetCommentRes{Data: []*pb.Comment{}}, nil
	}

	// 2. 获取主评论 ID 列表
	var ids []int32
	for _, v := range mainComments {
		ids = append(ids, v.Id)
	}

	// 3. 一次性查询所有主评论的前6条子评论（LIMIT 无法在每组控制，用 row_number 更合理，此处保守方式处理）
	// 替换第三步，使用原生 SQL 加 row_number 分组限制子评论条数
	var children []*model.VideoWorkComment

	pidList := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ids)), ","), "[]")

	sql := fmt.Sprintf(`
    SELECT *
    FROM (
        SELECT *,
               ROW_NUMBER() OVER (PARTITION BY pid ORDER BY created_at ASC) as rn
        FROM video_work_comment
        WHERE pid IN (%s)
    ) AS tmp
    WHERE tmp.rn <= 5
`, pidList)

	err = global.DB.Raw(sql).Scan(&children).Error
	if err != nil {
		return &pb.GetCommentRes{}, err
	}

	// 4. 将子评论按 parent_id 分组
	childrenMap := make(map[int32][]*model.VideoWorkComment)
	for _, child := range children {
		childrenMap[int32(child.Pid)] = append(childrenMap[int32(child.Pid)], child)
	}

	// 5. 构造返回结构
	var result []*pb.Comment
	for _, main := range mainComments {
		item := &pb.Comment{
			UserName: "",
			UserId:   0,
			Content:  main.Content,
			Avator:   "",
			Time:     FormatCommentTime(main.CreatedAt),
			Id:       main.Id,
			Data:     nil,
		}

		childList := childrenMap[main.Id]
		if len(childList) > 0 {
			// 最多展示 5 条
			max := 5
			if len(childList) < 5 {
				max = len(childList)
			}

			for i := 0; i < max; i++ {
				item.Data = append(item.Data, &pb.CommentP{
					UserName:  "",
					UserId:    int32(childList[i].UserId),
					Content:   childList[i].Content,
					Avator:    "",
					PUserId:   int32(childList[i].Pid),
					PUserName: "",
					Time:      FormatCommentTime(childList[i].CreatedAt),
				})
			}
		}

		// 设置是否有更多
		if int(main.ReplyCount) > len(item.Data) {
			item.HasMore = true
		} else {
			item.HasMore = false
		}

		result = append(result, item)
	}

	return &pb.GetCommentRes{
		Data: result,
	}, nil
}

func (c VideoWorksCommentServer) GetCommentReplyList(ctx context.Context, in *pb.GetCommentReplyListReq) (*pb.GetCommentReplyListRes, error) {
	fmt.Println(in)
	start := (in.Page - 1) * in.Limit
	var mainComments []*model.VideoWorkComment
	err := global.DB.Table("video_work_comment").Where("pid = ?", in.Pid).Offset(int(start)).Limit(int(in.Limit)).Find(&mainComments).Error
	if err != nil {
		return &pb.GetCommentReplyListRes{}, err
	}
	var ids []int32
	for _, v := range mainComments {
		ids = append(ids, int32(v.UserId))
		ids = append(ids, v.RepliedUserId)
	}

	var users []*model.VideoUser
	err = global.DB.Table("video_user").Where("id IN (?)", ids).Find(&users).Error
	if err != nil {
		return &pb.GetCommentReplyListRes{}, err
	}
	userMap := make(map[int32]*model.VideoUser)
	for _, v := range users {
		userMap[v.Id] = v
	}

	var children []*pb.CommentP
	var child *pb.CommentP
	for _, v := range mainComments {

		child = &pb.CommentP{
			UserName: userMap[int32(v.UserId)].Name,
			UserId:   int32(v.UserId),
			Content:  v.Content,
			Avator:   "",

			Time: FormatCommentTime(userMap[int32(v.UserId)].CreatedAt),
		}
		if v.RepliedUserId != 0 {
			child.PUserId = v.RepliedUserId
			child.PUserName = userMap[int32(v.RepliedUserId)].Name
		}

		children = append(children, child)
	}

	return &pb.GetCommentReplyListRes{Data: children}, nil
}
