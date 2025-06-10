package middlewear

import (
	"api/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		userId, err := pkg.VieyToken(token)
		if err != nil {
			fmt.Println(err)
		}

		// 设置 example 变量
		c.Set("userId", userId)

		// 请求前

		c.Next()

		// 请求后

	}
}
