package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"onlineVideo/models/dto"
	"onlineVideo/service"
	"onlineVideo/utils"
	"strconv"
)

// UserLogin 用户登录
func UserLogin(c *gin.Context) {
	var loginUser dto.UserLoginDto
	err := c.ShouldBind(&loginUser)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg:  utils.RequestError,
		})
		return
	}
	res := service.UserLogin(&loginUser)
	c.JSON(http.StatusOK, res)
}

// UserRegister 用户注册
func UserRegister(c *gin.Context) {
	var registerUser dto.UserRegisterDto
	err := c.ShouldBind(&registerUser)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg:  utils.RequestError,
		})
		return
	}
	res := service.UserRegister(registerUser.Username, registerUser.Password)
	c.JSON(http.StatusOK, res)
}

// CheckUserIsExit 验证用户名是否存在
func CheckUserIsExit(c *gin.Context) {
	var userName dto.UserNameDto
	err := c.ShouldBind(&userName)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg:  utils.RequestError,
		})
		return
	}
	res := service.CheckUserIsExist(userName.Username)
	c.JSON(http.StatusOK, res)
}

// UserInfo 获取用户信息
func UserInfo(c *gin.Context) {
	uid := c.GetUint("user_id")
	res := service.GetUserInfo(uid)
	c.JSON(http.StatusOK, res)
}

// GetUserInfo 根据id获取用户信息
func GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("uid"))
	res := service.GetUserInfo(uint(id))
	c.JSON(http.StatusOK, res)
}

// UpdateUserInfo 更新用户信息
func UpdateUserInfo(c *gin.Context) {
	uid := c.GetUint("user_id")
	var userData dto.UserInfoUpdateDto
	err := c.ShouldBind(&userData)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg:  utils.RequestError,
		})
		return
	}
	res := service.UpdateUserInfo(uid, &userData)
	c.JSON(http.StatusOK, res)
}

// UpdateUserPass 更新用户密码
func UpdateUserPass(c *gin.Context) {
	uid := c.GetUint("user_id")
	var userPass dto.UserUpdatePassDto
	err := c.ShouldBind(&userPass)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg:  utils.RequestError,
		})
		return
	}
	res := service.UpdateUserPass(uid, &userPass)
	c.JSON(http.StatusOK, res)
}
