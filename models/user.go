package models

type User struct {
	Id       uint `gorm:"primarykey"`
	Avatar   string
	Username string `gorm:"type:varchar(20);unique;not null"`
	Password string `gorm:"not null;"`
	Gender   uint   `gorm:"default:0"`
	Sign     string
}
