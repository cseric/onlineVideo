package service

import (
	"onlineVideo/common"
	"onlineVideo/models"
	"onlineVideo/utils"
)

// UploadAvatar 修改头像
func UploadAvatar(uid uint, avatarName string) utils.Response {
	DB := common.GetDB()
	DB.Model(&models.User{}).Where("id = ?", uid).Update("avatar", avatarName)
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg: utils.UploadSuccess,
	}
}
