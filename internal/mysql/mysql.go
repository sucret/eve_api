package mysql

import (
	"eve/internal/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

func GetConnection() (db *gorm.DB) {
	// 获取数据库配置
	config := global.Config.Database

	// 拼装dsn
	dsn := config.Username + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port +
		")/" + config.Database + "?charset=" + config.Charset + "&parseTime=True&loc=Local"

	// 初始化配置
	mysqlConfig := mysql.Config{
		DSN: dsn,
	}

	// 连接数据库
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用表名复数
		},
	})
	if err != nil {
		log.Println("数据库链接失败：", err)
		panic(err)
	}

	return
}
