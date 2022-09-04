package v1

import (
	"ginDemo/model/database"
	"ginDemo/model/response"
	"ginDemo/service"
	"ginDemo/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"regexp"
)

// Register
// @Summary 注册
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param data body response.RegisterQ true "用户名，密码，确认密码，邮箱"
// @Success 200 {object} response.CommonA
// @Router /register [post]
func Register(c *gin.Context) {
	var data response.RegisterQ
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusOK, response.CommonA{Message: "格式错误: " + err.Error(), Success: false})
		return
	}
	matched, _ := regexp.Match("^[A-Za-z\\d]{8,40}$", []byte(data.Password1))
	if !matched {
		c.JSON(http.StatusOK, response.CommonA{Message: "密码格式错误", Success: false})
		return
	}
	if data.Password1 != data.Password2 {
		c.JSON(http.StatusOK, response.CommonA{Message: "两次输入的密码不一致", Success: false})
		return
	}
	HashedPassword, _ := bcrypt.GenerateFromPassword([]byte(data.Password1), bcrypt.DefaultCost)
	user := database.User{Username: data.Username, Password: string(HashedPassword), Email: data.Email}
	if err := service.CreateUser(&user); err != nil {
		c.JSON(http.StatusOK, response.CommonA{Message: "注册失败", Success: false})
		return
	}
	c.JSON(http.StatusOK, response.CommonA{Message: "注册成功", Success: true})
}

// Login
// @Summary 登录
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param data body response.LoginQ true "用户名，密码"
// @Success 200 {object} response.LoginA
// @Router /login [post]
func Login(c *gin.Context) {
	var data response.LoginQ
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusOK, response.LoginA{CommonA: response.CommonA{Message: "格式错误: " + err.Error(), Success: false}})
		return
	}
	user, notFound := service.QueryUserByUsername(data.Username)
	if notFound {
		c.JSON(http.StatusOK, response.LoginA{CommonA: response.CommonA{Message: "用户不存在", Success: false}})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		c.JSON(http.StatusOK, response.LoginA{CommonA: response.CommonA{Message: "密码错误", Success: false}})
		return
	}
	token := utils.GenerateToken(user.UserID)
	c.JSON(http.StatusOK, response.LoginA{
		CommonA: response.CommonA{Message: "登录成功", Success: true},
		Token:   token,
		User:    user,
	})
}

// GetUserInfo
// @Summary 获取用户信息
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param data body response.GetUserInfoQ true "用户 ID"
// @Success 200 {object} response.GetUserInfoA
// @Security ApiKeyAuth
// @Router /user/info [post]
func GetUserInfo(c *gin.Context) {
	poster, _ := c.Get("user")
	var data response.GetUserInfoQ
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusOK, response.GetUserInfoA{CommonA: response.CommonA{Message: "格式错误: " + err.Error(), Success: false}})
		return
	}
	user, notFound := service.QueryUserByUserID(data.UserID)
	if notFound {
		c.JSON(http.StatusOK, response.GetUserInfoA{
			CommonA: response.CommonA{Message: "用户不存在", Success: false},
			Poster:  poster.(database.User),
		})
		return
	}
	c.JSON(http.StatusOK, response.GetUserInfoA{
		CommonA: response.CommonA{Message: "获取成功", Success: true},
		Poster:  poster.(database.User),
		User:    user,
	})
}
