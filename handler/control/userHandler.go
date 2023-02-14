package control

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"onlineVideo/models/dto"
	"onlineVideo/service"
	"onlineVideo/utils"
	"strconv"
)

// GetUserList 管理员获取用户列表
func GetUserList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "5"))
	if page <= 0 || pageSize <= 0 {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.PageOrPageSizeError,
		})
		return
	}
	res := service.AdminGetUserList(page, pageSize)
	c.JSON(http.StatusOK, res)
}

// SearchUser 管理员搜索用户
func SearchUser(c *gin.Context) {
	keyword := c.Query("keyword")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "5"))
	if len(keyword) == 0 {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.NoKeyword,
		})
		return
	}
	if page <= 0 || pageSize <= 0 {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.PageOrPageSizeError,
		})
		return
	}
	res := service.AdminSearchUser(keyword, page, pageSize)
	c.JSON(http.StatusOK, res)
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	var userId dto.AdminIdDto
	err := c.ShouldBind(&userId)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.RequestError,
		})
		return
	}
	res := service.AdminDeleteUser(userId.Id)
	c.JSON(http.StatusOK, res)
}
