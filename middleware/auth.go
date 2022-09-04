package middleware

import (
	"ginDemo/model/response"
	"ginDemo/service"
	"ginDemo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthRequired() gin.HandlerFunc { // Token 放在 Header 的 Authorization中, 并使用 Bearer 开头
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, response.CommonA{Message: "请求头为空", Success: false})
			c.Abort()
			return
		}
		id, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, response.CommonA{Message: "校验失败", Success: false})
			c.Abort()
			return
		}
		if user, notFound := service.QueryUserByUserID(id); notFound {
			c.JSON(http.StatusOK, response.CommonA{Message: "用户不存在", Success: false})
			c.Abort()
		} else {
			c.Set("user", user)
		}
	}
}
