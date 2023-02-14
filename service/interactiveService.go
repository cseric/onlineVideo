package service

import (
	"gorm.io/gorm"
	"onlineVideo/common"
	"onlineVideo/models"
	"onlineVideo/utils"
	"time"
)

// AddLike 点赞
func AddLike(uid uint, vid uint) utils.Response {
	DB := common.GetDB()
	status := isLike(DB, uid, vid)	// 验证是否点赞
	if status == 2 {	// 已点赞
		return utils.Response{
			Code: utils.CodeFail,
			Msg: "已经点过赞了",
		}
	}
	if status == 0 {	// 没有交互记录
		newLike := models.Interactive{
			Uid: uid,
			Vid: vid,
			IsLike: true,
		}
		DB.Create(&newLike)
	} else {	// 有记录, 更新状态
		DB.Model(&models.Interactive{}).Where("uid = ? and vid = ?", uid, vid).Update("is_like", true)
	}
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg: "点赞成功",
	}
}

// CancelLike 取消点赞
func CancelLike(uid uint, vid uint) utils.Response  {
	DB := common.GetDB()
	status := isLike(DB, uid, vid)	// 验证是否点赞
	if status == 2 {	// 已点赞，取消点赞
		DB.Model(&models.Interactive{}).Where("uid = ? and vid = ?", uid, vid).Update("is_like", false)
		return utils.Response{
			Code: utils.CodeSuccess,
			Msg: "取赞成功",
		}
	}
	return utils.Response{
		Code: utils.CodeFail,
		Msg: "还未点赞",
	}
}

// AddCollect 添加收藏
func AddCollect(uid uint, vid uint) utils.Response {
	DB := common.GetDB()
	status := isCollect(DB, uid, vid)	// 验证是否收藏
	if status == 2 {	// 已收藏
		return utils.Response{
			Code: utils.CodeFail,
			Msg: "已经收藏了",
		}
	}
	if status == 0 {	// 没有交互记录
		newCollect := models.Interactive{
			Uid: uid,
			Vid: vid,
			IsCollect: true,
			CollectTime: time.Now(),
		}
		DB.Create(&newCollect)
	} else {	// 有记录，更新状态
		DB.Model(&models.Interactive{}).Where("uid = ? and vid = ?", uid, vid).
			Updates(map[string]interface{}{
			"is_collect": true,
			"collect_time": time.Now(),
		})
	}
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg: "收藏成功",
	}
}

// CancelCollect 取消收藏
func CancelCollect(uid uint, vid uint) utils.Response {
	DB := common.GetDB()
	status := isCollect(DB, uid, vid)	// 验证是否收藏
	if status == 2 {	// 已收藏，取消收藏
		DB.Model(&models.Interactive{}).Where("uid = ? and vid = ?", uid, vid).Update("is_collect", false)
		return utils.Response{
			Code: utils.CodeSuccess,
			Msg: "取消收藏成功",
		}
	}
	return utils.Response{
		Code: utils.CodeFail,
		Msg: "还未收藏",
	}
}

// 验证是否点赞
func isLike(DB *gorm.DB, uid uint, vid uint) int {
	var like models.Interactive
	DB.Where("uid = ? and vid = ?", uid, vid).First(&like)
	if like.Id == 0 {	// 不存在
		return 0
	} else if !like.IsLike {	// 未点赞
		return 1
	}
	return 2	// 已点赞
}

// 验证是否收藏
func isCollect(DB *gorm.DB, uid uint, vid uint) int {
	var collect models.Interactive
	DB.Where("uid = ? and vid = ?", uid, vid).First(&collect)
	if collect.Id == 0 {	// 不存在
		return 0
	} else if !collect.IsCollect {	// 未点赞
		return 1
	}
	return 2	// 已点赞
}
