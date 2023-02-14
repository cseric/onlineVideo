package control

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onlineVideo/models/dto"
	"onlineVideo/service"
	"onlineVideo/utils"
	"strconv"
)

// GetCommentList 获取评论列表
func GetCommentList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "5"))
	if page <= 0 || pageSize <= 0 {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.PageOrPageSizeError,
		})
		return
	}
	res := service.AdminGetCommentList(page, pageSize)
	c.JSON(http.StatusOK, res)
}

// SearchComment 搜索评论
func SearchComment(c *gin.Context) {
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
	res := service.AdminSearchComment(keyword, page, pageSize)
	c.JSON(http.StatusOK, res)
}

// DeleteComment 删除评论
func DeleteComment(c *gin.Context) {
	var commentId dto.AdminIdDto
	err := c.ShouldBind(&commentId)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.RequestError,
		})
		return
	}
	res := service.AdminDeleteComment(commentId.Id)
	c.JSON(http.StatusOK, res)
}