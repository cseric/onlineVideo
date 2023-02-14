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

// --------------------------------用户--------------------------------

// UserLogin 用户登录
func UserLogin(loginUser *dto.UserLoginDto) utils.Response {
	DB := common.GetDB()
	// 查询用户是否存在
	var user models.User
	DB.Where("username = ?", loginUser.Username).First(&user)
	if user.Id == 0 {
		return utils.Response{
			Code: utils.CodeFail,
			Msg:  utils.UserNotExist,
		}
	}
	// 判断密码是否一致
	if loginUser.Password != user.Password {
		return utils.Response{
			Code: utils.CodeFail,
			Msg:  utils.PasswordError,
		}
	}
	// 发放token
	token, err := common.ReleaseToken(user.Id)
	if err != nil {
		return utils.Response{
			Code: utils.CodeServerError,
			Msg:  utils.ServerError,
		}
	}
	// 生成用户信息
	userInfo := vo.ToUserInfoVo(&user)
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg:  utils.LoginSuccess,
		Data: gin.H{
			"token":    common.TokenPrefix + token,
			"userInfo": userInfo,
		},
	}
}

// UserRegister 用户注册
func UserRegister(username string, password string) utils.Response {
	DB := common.GetDB()
	// 查询用户名是否存在
	if IsUserExist(DB, username) {
		return utils.Response{
			Code: utils.CodeFail,
			Msg:  utils.UserIsExist,
		}
	}
	// 添加用户
	newUser := models.User{
		Username: username,
		Password: password,
	}
	DB.Create(&newUser)
	// 返回结果
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg:  utils.RegisterSuccess,
	}
}

// CheckUserIsExist 验证用户名是否存在（注册时）
func CheckUserIsExist(username string) utils.Response {
	DB := common.GetDB()
	isExist := IsUserExist(DB, username)
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg:  utils.GetDataSuccess,
		Data: gin.H{
			"isExist": isExist,
		},
	}
}

// GetUserInfo 获取用户信息
func GetUserInfo(uid uint) utils.Response {
	var user models.User
	DB := common.GetDB()
	DB.First(&user, uid)
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg:  utils.GetDataSuccess,
		Data: gin.H{
			"userInfo": vo.ToUserInfoVo(&user),
		},
	}
}

// UpdateUserInfo 更新用户信息
func UpdateUserInfo(uid uint, userInfo *dto.UserInfoUpdateDto) utils.Response {
	DB := common.GetDB()
	// 查询用户名是否存在
	var user models.User
	DB.Where("username = ? and id != ?", userInfo.Username, uid).First(&user)
	if user.Id != 0 {
		return utils.Response{
			Code: utils.CodeFail,
			Msg:  utils.UserIsExist,
		}
	}
	// 更新用户信息
	DB.Model(&models.User{}).Where("id = ?", uid).
		Updates(map[string]interface{}{"username": userInfo.Username, "gender": userInfo.Gender, "sign": userInfo.Sign})
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg:  utils.UpdateSuccess,
	}
}

// UpdateUserPass 修改用户密码
func UpdateUserPass(uid uint, userPass *dto.UserUpdatePassDto) utils.Response {
	DB := common.GetDB()
	var user models.User
	DB.Select("password").Where("id = ?", uid).First(&user)
	if userPass.OldPassword != user.Password { // 原密码输入错误
		return utils.Response{
			Code: utils.CodeFail,
			Msg:  utils.PasswordError,
		}
	}
	// 更新用户密码
	DB.Model(&models.User{}).Select("password").Where("id = ?", uid).Updates(map[string]interface{}{
		"password": userPass.NewPassword,
	})
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg:  utils.UpdateSuccess,
	}
}

// --------------------------------管理员--------------------------------

// AdminGetUserList 管理员获取用户列表
func AdminGetUserList(page int, pageSize int) utils.Response {
	var users []vo.UserInfoVo
	DB := common.GetDB()

	// 获取用户总数
	var total int64
	DB.Model(&models.User{}).Count(&total)
	// 分页查询
	DB = DB.Limit(pageSize).Offset((page - 1) * pageSize)
	DB.Model(&models.User{}).Scan(&users)
	// 查询成功
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg:  utils.GetDataSuccess,
		Data: gin.H{
			"total": total,
			"users": users,
		},
	}
}

// AdminSearchUser 管理员搜索用户
func AdminSearchUser(keyword string, page int, pageSize int) utils.Response {
	var total int64
	var users []vo.UserInfoVo

	DB := common.GetDB()
	keyword = "%" + keyword + "%"
	DB.Model(&models.User{}).Where("id like ? or username like ?", keyword, keyword).Count(&total)
	DB = DB.Limit(pageSize).Offset((page - 1) * pageSize)
	DB.Model(&models.User{}).Where("id like ? or username like ?", keyword, keyword).Scan(&users)
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg:  utils.SearchSuccess,
		Data: gin.H{
			"total": total,
			"users": users,
		},
	}
}

// AdminDeleteUser 管理员删除用户
func AdminDeleteUser(userId uint) utils.Response {
	DB := common.GetDB()
	DB.Delete(&models.User{}, userId)                                       // 删除用户
	DB.Where("uid = ?", userId).Delete(&models.Video{})                     // 删除用户视频
	DB.Where("uid = ?", userId).Delete(&models.Comment{})                   // 删除用户评论
	DB.Where("uid = ?", userId).Delete(&models.Interactive{})               // 删除用户交互数据
	DB.Where("uid = ? or fid = ?", userId, userId).Delete(&models.Follow{}) // 删除用户关系
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg:  utils.DeleteSuccess,
	}
}

// --------------------------------功能性函数--------------------------------

// IsUserExist 判断用户名是否存在
func IsUserExist(DB *gorm.DB, username string) bool {
	var user models.User
	DB.Where("username = ?", username).First(&user)
	if user.Id != 0 {
		return true
	}
	return false
}
