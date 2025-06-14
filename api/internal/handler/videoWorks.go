package handler

import (
	"api/internal/server"
	"api/middlewear"
	pb "api/proto/videoWorks"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"strings"
)

type GetVideoWorksReq struct {
	Page int `json:"page" form:"page"`
}

func GetVideoWorks(c *gin.Context) {
	var req GetVideoWorksReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	data, err := server.GetVideoWorks(c, &pb.GetVideoWorksReq{
		Page: int32(req.Page),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
		"code": 200,
	})
	return
}

type AddVideoWorksResp struct {
	Title    string `form:"title" json:"title"`
	Desc     string `form:"desc" json:"desc"`
	Music    int32  `form:"music" json:"music_id"`
	WorkType string `form:"work_type" json:"work_type"`
}

func AddVideoWorks(c *gin.Context) {
	var req AddVideoWorksResp
	userId := c.MustGet("userId").(int32)
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Invalid request data: " + err.Error(),
		})
		return
	}

	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Failed to get uploaded file: " + err.Error(),
		})
		return
	}
	if file.Size > 1024*1024*50 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "文件内容过大",
		})
	}

	// 验证文件扩展名
	allowedExtensions := map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
		".mp4":  true,
	}

	ext := strings.ToLower(path.Ext(file.Filename))
	if !allowedExtensions[ext] {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "文件类型不支持，仅支持 png, jpg, jpeg, mp4 格式",
		})
		return
	}

	// 直接上传到 MinIO
	fileUrl, err := middlewear.UploadFileToMinIO("lanjing", file.Filename, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "Failed to upload file to MinIO: " + err.Error(),
		})
		return
	}

	data := &pb.AddVideoWorksReq{
		UserId:   userId,
		Title:    req.Title,
		Desc:     req.Desc,
		MusicId:  req.Music,
		WorkType: req.WorkType,
		Ip:       c.ClientIP(),
		FilePath: fileUrl,
	}

	res, err := server.AddVideoWorks(c, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": res,
		"code": 200,
	})
	return

}
