package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"onlineVideo/common"
	"onlineVideo/models"
	"onlineVideo/models/dto"
	"onlineVideo/models/vo"
	"onlineVideo/utils"
)

// --------------------------------管理员公共接口--------------------------------

// AdminLogin 管理员登录
func AdminLogin(loginAdmin *dto.AdminLoginDto) utils.Response {
	res := utils.Response{
		Code: utils.CodeSuccess,
		Msg:  utils.LoginSuccess,
	}

	var admin models.Admin
	DB := common.GetDB()
	DB.Where("username = ?", loginAdmin.Username).First(&admin)
	// 用户不存在
	if admin.Id == 0 {
		res.Code = utils.CodeFail
		res.Msg = utils.UserNotExist
		return res
	}

	// 密码错误
	if loginAdmin.Password != admin.Password {
		res.Code = utils.CodeFail
		res.Msg = utils.PasswordError
		return res
	}

	// 发放token
	token, err := common.ReleaseToken(admin.Id)
	if err != nil { // 服务错误
		res.Code = utils.CodeServerError
		res.Msg = utils.ServerError
		return res
	}

	// 保存管理员信息
	var adminInfo vo.AdminInfoVo
	adminInfo.Id = admin.Id
	adminInfo.Username = admin.Username
	if admin.Authority == models.SuperAdminAuth {
		adminInfo.Role = "超级管理员"
	} else if admin.Authority == models.AdminAuth {
		adminInfo.Role = "管理员"
	} else if admin.Authority == models.AuditAuth {
		adminInfo.Role = "审核员"
	}
	res.Data = gin.H{
		"token":     common.TokenPrefix + token,
		"adminInfo": adminInfo,
	}
	return res
}

// UpdateAdminPassword 修改管理员密码
func UpdateAdminPassword(id uint, adminPass *dto.AdminPasswordDto) utils.Response {
	DB := common.GetDB()
	var admin models.Admin
	DB.Select("password").Where("id = ?", id).First(&admin)
	// 原密码错误
	if admin.Password != adminPass.OldPassword {
		return utils.Response{
			Code: utils.CodeFail,
			Msg:  utils.PasswordError,
		}
	}
	// 修改密码
	DB.Model(&models.Admin{}).Select("password").Where("id = ?", id).Updates(map[string]interface{}{"password": adminPass.NewPassword})
	// 修改成功
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg:  utils.UpdateSuccess,
	}
}

// GetStatisticsData 获取统计数据
func GetStatisticsData() utils.Response {
	DB := common.GetDB()
	var (
		adminCount int64 = 0
		userCount int64 = 0
		videoCount int64 = 0
		auditCount int64 = 0
		commentCount int64 = 0
	)
	DB.Model(&models.Admin{}).Where("authority < ?", models.SuperAdminAuth).Count(&adminCount)
	DB.Model(&models.User{}).Count(&userCount)
	DB.Model(&models.Video{}).Where("status = ?", models.CodeAuditPass).Count(&videoCount)
	DB.Model(&models.Video{}).Where("status = ?", models.CodeNotAudit).Count(&auditCount)
	DB.Model(&models.Comment{}).Count(&commentCount)
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg: utils.GetDataSuccess,
		Data: gin.H{
			"admins": adminCount,
			"users": userCount,
			"videos": videoCount,
			"audits": auditCount,
			"comments": commentCount,
		},
	}
}

// --------------------------------超级管理员--------------------------------

// GetAdminList 获取管理员列表
func GetAdminList(page int, pageSize int) utils.Response {
	var admins []vo.AdminListVo
	var total int64
	DB := common.GetDB()
	DB.Model(&models.Admin{}).Where("authority < 3").Count(&total)
	// 分页查询
	DB = DB.Limit(pageSize).Offset((page - 1) * pageSize)
	DB.Model(&models.Admin{}).Select("id, username, authority").Where("authority < 3").Find(&admins)
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg:  utils.GetDataSuccess,
		Data: gin.H{
			"admins": admins,
			"total": total,
		},
	}
}

// AddAdmin 添加管理员
func AddAdmin(admin *dto.AddAdminDto) utils.Response {
	DB := common.GetDB()
	if isAdminNameExist(DB, admin.Username) {
		return utils.Response{
			Code: utils.CodeFail,
			Msg:  utils.UserIsExist,
		}
	}
	newAdmin := models.Admin{
		Username:  admin.Username,
		Password:  admin.Password,
		Authority: admin.Authority,
	}
	DB.Create(&newAdmin)
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg:  utils.AddSuccess,
	}
}

// UpdateAdmin 修改管理员
func UpdateAdmin(admin *dto.UpdateAdminDto) utils.Response {
	DB := common.GetDB()
	DB.Model(&models.Admin{}).Where("id = ?", admin.Id).Update("authority", admin.Authority)
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg:  utils.UpdateSuccess,
	}
}

// DeleteAdmin 删除管理员
func DeleteAdmin(id uint) utils.Response {
	DB := common.GetDB()
	DB.Delete(&models.Admin{}, id)
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg:  utils.DeleteSuccess,
	}
}

// --------------------------------功能性函数--------------------------------

// isAdminNameExist 管理员用户名是否存在
func isAdminNameExist(DB *gorm.DB, username string) bool {
	var admin models.Admin
	DB.Model(&models.Admin{}).Where("username = ?", username).First(&admin)
	if admin.Id == 0 {
		return false
	}
	return true
}
