package core

import (
	"fmt"
	"go.uber.org/zap"
	"time"
	"yuyu/global"
	"yuyu/initialize"
	"yuyu/service/system"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.GvaConfig.System.UseMultipoint || global.GvaConfig.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	// 从db加载jwt数据
	if global.GvaDb != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.GvaConfig.System.Addr)
	s := initServer(address, Router)

	time.Sleep(10 * time.Microsecond)

	global.GvaLog.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 yu-miaomiao-server
	当前版本:v1.0.0
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	`, address)

	global.GvaLog.Error(s.ListenAndServe().Error())
}
