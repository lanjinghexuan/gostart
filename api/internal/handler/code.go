package handler

import (
	"api/internal/server"
	pb "api/proto/Code"
	"errors"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"regexp"
)

type SendCodeReq struct {
	Mobile string `json:"mobile" form:"mobile" binding:"required"`
}

func SendCode(c *gin.Context) {
	var req SendCodeReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pattern := `^1[3-9]\d{9}$`
	re := regexp.MustCompile(pattern)
	if !re.MatchString(req.Mobile) {
		c.JSON(http.StatusOK, gin.H{"code": 601, "data": errors.New("不是中国地区的手机号")})
		return
	}

	code := int32(rand.Intn(8999) + 1000)
	codereq := &pb.SendCodeReq{
		Mobile: req.Mobile,
		Code:   code,
	}

	res, err := server.SendCode(c, codereq)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 601, "data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": res})
	return
}
