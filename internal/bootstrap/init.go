package bootstrap

import (
	"eve/internal/config"
	"eve/internal/global"
)

func init() {
	// 初始化配置文件
	global.Config = config.GetConfig()
}
