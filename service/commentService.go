package service

import (
	"github.com/gin-gonic/gin"
	"onlineVideo/common"
	"onlineVideo/models"
	"onlineVideo/models/dto"
	"onlineVideo/models/vo"
	"onlineVideo/utils"
	"time"
)

// --------------------------------用户--------------------------------

// GetCommentByVid 用户通过视频id获取评论
func GetCommentByVid(vid uint, page int, pageSize int) utils.Response {
	var total int64
	var comments []vo.VideoCommentVo

	DB := common.GetDB()
	DB.Model(&models.Comment{}).Where("comment.vid = ?", vid).Count(&total)
	DB = DB.Limit(pageSize).Offset((page - 1) * pageSize)
	DB.Model(&models.Comment{}).
		Select("comment.id, comment.uid, comment.content, comment.comment_time, user.username, user.avatar").
		Joins("join user on user.id = comment.uid").Where("comment.vid = ?", vid).Find(&comments)
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg: utils.GetDataSuccess,
		Data: gin.H{
			"total": total,
			"comments": comments,
		},
	}
}

// DeleteCommentById 删除评论
func DeleteCommentById(cid uint) utils.Response {
	DB := common.GetDB()
	DB.Delete(&models.Comment{}, cid)
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg: utils.DeleteSuccess,
	}
}

// AddComment 添加评论
func AddComment(uid uint, comment *dto.CommentDto) utils.Response {
	DB := common.GetDB()
	commentTime := time.Now()	// 评论时间
	newComment := models.Comment{
		Vid: comment.Vid,
		Uid: uid,
		Content: comment.Content,
		CommentTime: commentTime,
	}
	DB.Create(&newComment)
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg: utils.CommentSuccess,
	}
}

// --------------------------------管理员--------------------------------

// AdminGetCommentList 管理员获取评论列表
func AdminGetCommentList(page int, pageSize int) utils.Response {
	var total int64
	var comments []vo.CommentInfoVo

	DB := common.GetDB()
	DB.Model(&models.Comment{}).Count(&total)
	DB = DB.Limit(pageSize).Offset((page - 1) * pageSize)
	DB.Model(&models.Comment{}).Select("comment.*, video.title, user.username").
		Joins("join video on video.id = comment.vid").Joins("join user on user.id = comment.uid").Find(&comments)
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg: utils.GetDataSuccess,
		Data: gin.H{
			"total":total,
			"comments": comments,
		},
	}
}

// AdminSearchComment 管理员搜索评论
func AdminSearchComment(keyword string, page int, pageSize int) utils.Response {
	var total int64
	var comments []vo.CommentInfoVo

	DB := common.GetDB()
	DB.Model(&models.Comment{}).Joins("join video on video.id = comment.vid").
		Joins("join user on user.id = comment.uid").
		Where("comment.uid like ? or comment.vid like ? or user.username like ?", keyword, keyword, keyword).Count(&total)
	DB = DB.Limit(pageSize).Offset((page - 1) * pageSize)
	keyword = "%" + keyword + "%"
	DB.Model(&models.Comment{}).Select("comment.*, video.title, user.username").
		Joins("join video on video.id = comment.vid").
		Joins("join user on user.id = comment.uid").
		Where("comment.uid like ? or comment.vid like ? or user.username like ?", keyword, keyword, keyword).Find(&comments)

	return utils.Response{
		Code: utils.CodeSuccess,
		Msg: utils.SearchSuccess,
		Data: gin.H{
			"total": total,
			"comments": comments,
		},
	}
}

// AdminDeleteComment 管理员删除评论
func AdminDeleteComment(commentId uint) utils.Response {
	DB := common.GetDB()
	DB.Delete(&models.Comment{}, commentId)

	return utils.Response{
		Code: utils.CodeSuccess,
		Msg: utils.DeleteSuccess,
	}
}
