package bootstrap

import (
	"eve/app/task"
	"eve/internal/config"
	"eve/internal/crontab"
	"eve/internal/event"
	"eve/internal/global"
	"eve/internal/logger"
	"eve/internal/mysql"
	"eve/internal/validator"
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

	// 注册验证器
	validator.InitValidator()

	// 初始化事件机制
	global.EventDispatcher = event.New()

	// 初始化定时任务
	global.Crontab = crontab.Init()
	global.Crontab.AddTask(task.Tasks()...)
	global.Crontab.Start()
}
