package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onlineVideo/models/dto"
	"onlineVideo/service"
	"onlineVideo/utils"
)

// AddLike 点赞
func AddLike(c *gin.Context) {
	var interactiveId dto.InteractiveIdDto
	err := c.ShouldBind(&interactiveId)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.RequestError,
		})
		return
	}
	if interactiveId.Id <= 0 {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.VideoNotExist,
		})
		return
	}
	uid := c.GetUint("user_id")
	res := service.AddLike(uid, interactiveId.Id)
	c.JSON(http.StatusOK, res)
}

// CancelLike 取消点赞
func CancelLike(c *gin.Context) {
	var interactiveId dto.InteractiveIdDto
	err := c.ShouldBind(&interactiveId)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.RequestError,
		})
		return
	}
	uid := c.GetUint("user_id")
	res := service.CancelLike(uid, interactiveId.Id)
	c.JSON(http.StatusOK, res)
}

// AddCollect 收藏
func AddCollect(c *gin.Context) {
	var interactiveId dto.InteractiveIdDto
	err := c.ShouldBind(&interactiveId)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.RequestError,
		})
		return
	}
	if interactiveId.Id <= 0 {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.VideoNotExist,
		})
		return
	}
	uid := c.GetUint("user_id")
	res := service.AddCollect(uid, interactiveId.Id)
	c.JSON(http.StatusOK, res)
}

// CancelCollect 取消收藏
func CancelCollect(c *gin.Context) {
	var interactiveId dto.InteractiveIdDto
	err := c.ShouldBind(&interactiveId)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.RequestError,
		})
		return
	}
	uid := c.GetUint("user_id")
	res := service.CancelCollect(uid, interactiveId.Id)
	c.JSON(http.StatusOK, res)
}
