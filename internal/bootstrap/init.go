package bootstrap

import (
	"eve/internal/config"
	"eve/internal/global"
	"eve/internal/mysql"
)

func init() {
	// 初始化配置文件
	global.Config = config.GetConfig()

	// 初始化数据库
	global.Db = mysql.GetConnection()
}
