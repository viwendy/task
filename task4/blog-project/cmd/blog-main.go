package main

import (
	"blog-main/config"
	"blog-main/global"
	"blog-main/router"
	"fmt"
	"net/http"
	"os"
)

func main() {
	config.InitConfig("../conf/config.yaml")
	global.InitMysql()
	global.InitRedis()
	r := router.InitRouter()
	addr := fmt.Sprintf("%s:%d", "localhost", config.AppConfig.Server.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("服务器启动失败：", err)
		os.Exit(-1)
	}
	fmt.Println("服务器启动")
}
