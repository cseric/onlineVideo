package models

type Follow struct {
	Id  uint `gorm:"primarykey"`
	Uid uint `gorm:"not null;"` // 用户id
	Fid uint `gorm:"not null;"` // 关注id
}
