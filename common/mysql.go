package common

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"onlineVideo/models"
)

var db *gorm.DB

func InitMySQL() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.pass"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.name"))
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
		DefaultStringSize: 255,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic("数据库连接失败，Error: " + err.Error())
	}

	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(10)

	db.AutoMigrate(&models.Admin{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Video{})
	db.AutoMigrate(&models.Comment{})
	db.AutoMigrate(&models.Follow{})
	db.AutoMigrate(&models.Interactive{})
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}
