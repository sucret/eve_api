package bootstrap

import (
	"eve/internal/config"
	"eve/internal/global"
	"eve/internal/logger"
	"eve/internal/mysql"
)

func init() {
	var err error

	// 初始化配置文件
	global.Config = config.GetConfig()

	// 初始化数据库
	global.DB = mysql.GetConnection()

	// 初始化日志
	if global.Logger, err = logger.New(); err != nil {
		panic(err)
		return
	}
}
