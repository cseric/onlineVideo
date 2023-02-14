package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"onlineVideo/common"
	"onlineVideo/models"
	"onlineVideo/models/vo"
	"onlineVideo/utils"
)

// AddFollow 添加关注
func AddFollow(uid uint, fid uint) utils.Response {
	DB := common.GetDB()
	// 关注的用户是否存在
	var user models.User
	DB.First(&user, fid)
	if user.Id == 0 {
		return utils.Response{
			Code: utils.CodeFail,
			Msg:  utils.UserNotExist,
		}
	}
	// 查询是否关注
	var follow models.Follow
	DB.Where("uid = ? and fid = ?", uid, fid).First(&follow)
	if follow.Id != 0 {
		return utils.Response{
			Code: utils.CodeFail,
			Msg:  "已经关注过了",
		}
	}
	// 添加关注记录
	newFollow := models.Follow{
		Uid: uid,
		Fid: fid,
	}
	DB.Create(&newFollow)
	// 返回结果
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg:  utils.FollowSuccess,
	}
}

// CancelFollow 取消关注
func CancelFollow(uid uint, fid uint) utils.Response {
	DB := common.GetDB()
	// 查询关注状态
	var follow models.Follow
	DB.Where("uid = ? and fid = ?", uid, fid).First(&follow)
	if follow.Id == 0 {
		return utils.Response{
			Code: utils.CodeFail,
			Msg:  "还没有关注哦",
		}
	}
	// 删除关注记录
	DB.Where("uid = ? and fid = ?", uid, fid).Delete(&models.Follow{})
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg:  utils.UnFollowSuccess,
	}
}

// GetFollowStatus 获取关注状态
func GetFollowStatus(uid uint, fid uint) utils.Response {
	DB := common.GetDB()
	status := isFollow(DB, uid, fid)
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg:  utils.GetDataSuccess,
		Data: gin.H{
			"follow": status,
		},
	}
}

// GetFollowData 获取关注和粉丝人数
func GetFollowData(uid uint) utils.Response {
	var follow int64 // 关注人数
	var fans int64   // 粉丝人数
	DB := common.GetDB()
	DB.Model(&models.Follow{}).Where("uid = ?", uid).Count(&follow)
	DB.Model(&models.Follow{}).Where("fid = ?", uid).Count(&fans)
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg:  utils.GetDataSuccess,
		Data: gin.H{
			"follow": follow,
			"fans":   fans,
		},
	}
}

// GetFollowList 获取关注列表
func GetFollowList(uid uint, page int, pageSize int) utils.Response {
	var follow []vo.FollowVo
	var total int64
	DB := common.GetDB()
	DB.Raw("select count(*) as total from `user` where id in (select fid from follow where uid = ?)", uid).Count(&total)
	DB = DB.Limit(pageSize).Offset((page - 1) * pageSize)
	DB.Raw("select id, username, avatar, sign from `user` where id in (select fid from follow where uid = ?)", uid).Find(&follow)
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg:  utils.GetDataSuccess,
		Data: gin.H{
			"total":       total,
			"followUsers": follow,
		},
	}
}

// GetFansList 获取粉丝列表
func GetFansList(uid uint, page int, pageSize int) utils.Response {
	var fans []vo.FollowVo
	var total int64
	DB := common.GetDB()
	DB.Raw("select count(*) as total from `user` where id in (select uid from follow where fid = ?)", uid).Count(&total)
	DB = DB.Limit(pageSize).Offset((page - 1) * pageSize)
	DB.Raw("select id, username, avatar, sign from `user` where id in (select uid from follow where fid = ?)", uid).Find(&fans)
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg:  utils.GetDataSuccess,
		Data: gin.H{
			"total":     total,
			"fansUsers": fans,
		},
	}
}

// isFollow 获取关注状态
func isFollow(DB *gorm.DB, uid uint, fid uint) bool {
	var follow models.Follow
	DB.Where("uid = ? and fid = ?", uid, fid).First(&follow)
	if follow.Id != 0 {
		return true
	}
	return false
}
