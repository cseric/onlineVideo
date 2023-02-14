package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onlineVideo/models/dto"
	"onlineVideo/service"
	"onlineVideo/utils"
	"strconv"
)

// GetVideoComment 获取视频评论
func GetVideoComment(c *gin.Context) {
	vid, _ := strconv.Atoi(c.Query("vid"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "5"))
	if page <= 0 || pageSize <= 0 {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.PageOrPageSizeError,
		})
		return
	}
	if vid <= 0 {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.VideoNotExist,
		})
		return
	}
	res := service.GetCommentByVid(uint(vid), page, pageSize)
	c.JSON(http.StatusOK, res)
}

// DeleteVideoComment 删除视频评论
func DeleteVideoComment(c *gin.Context) {
	var commentId dto.CommentIdDto
	err := c.ShouldBind(&commentId)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.RequestError,
		})
		return
	}
	res := service.DeleteCommentById(commentId.Id)
	c.JSON(http.StatusOK, res)
}

// AddComment 添加评论
func AddComment(c *gin.Context) {
	var comment dto.CommentDto
	err := c.ShouldBind(&comment)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.RequestError,
		})
		return
	}
	uid := c.GetUint("user_id")
	res := service.AddComment(uid, &comment)
	c.JSON(http.StatusOK, res)
}
