package vo

import "onlineVideo/models"

// UserInfoVo 用户信息
type UserInfoVo struct {
	Id       uint   `json:"id"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	Gender   uint   `json:"gender"`
	Sign     string `json:"sign"`
}

// ToUserInfoVo 转换用户信息格式
func ToUserInfoVo(user *models.User) UserInfoVo {
	return UserInfoVo{
		Id:       user.Id,
		Avatar:   user.Avatar,
		Username: user.Username,
		Gender:   user.Gender,
		Sign:     user.Sign,
	}
}
