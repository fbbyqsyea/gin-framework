package initialize

import (
	"fmt"

	"github.com/fbbyqsyea/gin-framework/global"
	"github.com/spf13/viper"
)

// 初始化配置
func initConfig() {
	v := viper.New()
	// 配置文件path
	v.SetConfigFile("config.yaml")
	// 配置文件格式
	v.SetConfigType("yaml")
	// 读取配置文件
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error read config: [%s]", err))
	}
	// 格式化到global.CONFIG
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		panic(fmt.Errorf("fatal error unmarshal config: [%s]", err))
	}
}
