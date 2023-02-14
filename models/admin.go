package models

type Admin struct {
	Id        uint   `gorm:"primarykey"`
	Username  string `gorm:"type:varchar(20);unique;not null"`
	Password  string `gorm:"not null"`
	Authority uint   `gorm:"not null"`
}

const (
	SuperAdminAuth uint = 3 // 超级管理员
	AdminAuth      uint = 2 // 管理员
	AuditAuth      uint = 1 // 审核员
	PublicAuth     uint = 0
)
