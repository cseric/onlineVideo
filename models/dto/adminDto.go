package dto

type AdminIdDto struct {
	Id uint `json:"id" binding:"required"`
}

type AdminLoginDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateAdminDto struct {
	Id        uint   `json:"id" binding:"required"`
	Authority uint `json:"authority" binding:"required"`
}

type AddAdminDto struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Authority uint   `json:"authority" binding:"required"`
}

type AdminPasswordDto struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
	RePassword  string `json:"re_password" binding:"required,eqfield=NewPassword"`
}
