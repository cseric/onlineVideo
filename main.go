package main

import (
	"github.com/spf13/viper"
	"onlineVideo/common"
	"onlineVideo/conf"
	"onlineVideo/routes"
)

func main() {
	// 初始化配置
	conf.InitConfig()

	// 初始化数据库
	common.InitMySQL()
	defer common.CloseDB()

	// 加载路由
	router := routes.SetupRoutes()

	// 启动服务
	if err := router.Run(":"+viper.GetString("app.server_port")); err != nil {
		panic("服务启动失败，Error: " + err.Error())
	}
}
