package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"web/global"
	"web/initialize"
)

func main() {
	initialize.InitLogger()

	initialize.InitConfig()

	initialize.InitMysql()

	g := initialize.InitRouter()

	serverConfig := global.ServerConfig

	fmt.Println("服务启动中,端口:", serverConfig.Port)

	go func() {
		err := g.Run(fmt.Sprintf(":%d", serverConfig.Port))
		if err != nil {
			panic("启动失败:" + err.Error())
		}
	}()

	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
