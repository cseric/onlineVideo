package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("config")		// 配置文件名（不带扩展名）
	viper.SetConfigType("yaml")			// 配置文件类型
	viper.AddConfigPath("./conf/")		// 配置文件路径

	err := viper.ReadInConfig()
	if err != nil {
		panic("初始化配置失败，Error: " + err.Error())
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置已修改！")
	})
}
