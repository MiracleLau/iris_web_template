package main

import (
	"fmt"
	"web_template/app"
	"web_template/config"
	"web_template/db"
	"web_template/utils"
)

func main(){
	fmt.Println("读取配置文件")

	config.GetConfig("config/config.yaml")

	fmt.Println("初始化日志")

	utils.InitLogger(config.Config.LogPath)

	app.InitApp()

	app.InitRoute()

	app.SetApplicationLog(config.Config.LogPath)

	db.Connect()

	db.AutoMigrate()

	app.Start()
}