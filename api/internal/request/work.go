package request

type WorkAdd struct {
	Title          string `form:"title" binding:"required"`           //标题
	Describe       string `form:"describe" binding:"required"`        //描述
	MusicId        int    `form:"music_id" binding:"required"`        //选择音乐
	Status         string `form:"status" binding:"required"`          //审核状态
	WorkCate       string `form:"work_cate" binding:"required"`       //作品类型
	Reviewer       string `form:"reviewer" binding:"required"`        //审核人
	WorkPermission string `form:"work_permission" binding:"required"` //作品权限
	Address        string `form:"address" binding:"required"`         //ip地址
	LikeNum        int    `form:"like_num" binding:"required"`        //喜欢数量
	ContentNum     int    `form:"content_num" binding:"required"`     //评论数
	ShareNum       int    `form:"share_num" binding:"required"`       //分享数
	CollectNum     int    `form:"collect_num" binding:"required"`     //收藏数
	BrowseNum      int    `form:"browse_num" binding:"required"`      //浏览量
}

type WorkStatus struct {
	WorkCate string `form:"work_cate" binding:"required"` //作品类型
}
