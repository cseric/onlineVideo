package dto

type UserLoginDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRegisterDto struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required,min=6"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type UserNameDto struct {
	Username string `json:"username" binding:"required"`
}

type UserInfoUpdateDto struct {
	Username string `json:"username" binding:"required"`
	Gender   int    `json:"gender" binding:"oneof=0 1 2"`
	Sign     string `json:"sign"`
}

type UserUpdatePassDto struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
	RePassword  string `json:"re_password" binding:"required,eqfield=NewPassword"`
}
