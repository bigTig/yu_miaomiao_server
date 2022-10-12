package main

import (
	"go.uber.org/zap"
	"yuyu/core"
	"yuyu/global"
	"yuyu/initialize"
)

// @title                       Swagger Example API
// @version                     0.0.1
// @description                 This is a sample Server pets
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	global.GvaVp = core.Viper() // 初始化Viper
	global.GvaLog = core.Zap()  // 初始化zap日志库
	zap.ReplaceGlobals(global.GvaLog)

	global.GvaDb = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()

	if global.GvaDb != nil {
		initialize.RegisterTables(global.GvaDb) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GvaDb.DB()
		defer db.Close()
	}

	core.RunWindowsServer()
}
