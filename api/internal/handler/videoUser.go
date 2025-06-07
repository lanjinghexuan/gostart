package handler

import "github.com/gin-gonic/gin"

type LoginReq struct {
	Mobile string `json:"mobile"`
}

func Login(c *gin.Context) {

}
