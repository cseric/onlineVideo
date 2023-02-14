package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onlineVideo/models/dto"
	"onlineVideo/service"
	"onlineVideo/utils"
	"strconv"
)

// AddFollow 关注
func AddFollow(c *gin.Context) {
	var followId dto.FollowIdDto
	err := c.ShouldBind(&followId)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.RequestError,
		})
		return
	}
	uid := c.GetUint("user_id")
	if followId.Id == uid {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: "不能对自己进行操作",
		})
		return
	}
	res := service.AddFollow(uid, followId.Id)
	c.JSON(http.StatusOK, res)
}

// CancelFollow 取消关注
func CancelFollow(c *gin.Context) {
	var followId dto.FollowIdDto
	err := c.ShouldBind(&followId)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.RequestError,
		})
		return
	}
	uid := c.GetUint("user_id")
	res := service.CancelFollow(uid, followId.Id)
	c.JSON(http.StatusOK, res)
}

// GetFollowStatus 获取关注状态
func GetFollowStatus(c *gin.Context) {
	fid, _ := strconv.Atoi(c.Query("fid"))
	if fid == 0 {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.UserNotExist,
		})
		return
	}
	uid := c.GetUint("user_id")
	res := service.GetFollowStatus(uid, uint(fid))
	c.JSON(http.StatusOK, res)
}

// GetFollowData 获取关注和粉丝数据
func GetFollowData(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Query("uid"))
	if uid == 0 {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.UserNotExist,
		})
		return
	}
	res := service.GetFollowData(uint(uid))
	c.JSON(http.StatusOK, res)
}

// GetFollowList 获取关注列表
func GetFollowList(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Query("uid"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "5"))
	if uid == 0 {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.UserNotExist,
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
	res := service.GetFollowList(uint(uid), page, pageSize)
	c.JSON(http.StatusOK, res)
}

// GetFansList 获取粉丝列表
func GetFansList(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Query("uid"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "5"))
	if uid == 0 {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.UserNotExist,
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
	res := service.GetFansList(uint(uid), page, pageSize)
	c.JSON(http.StatusOK, res)
}
