package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/config"
	"github.com/zheng-yi-yi/simpledouyin/setup"
)

func main() {
	Init()             // 初始化所有配置。
	r := gin.Default() // 创建一个 HTTP 服务器的实例，并使用默认的配置选项进行初始化。
	initRouter(r)      // 初始化路由规则并将这些规则添加到一个名为 r 的 Gin 路由引擎中。
	r.Run()            // 启动 Gin 服务器，开始监听来自客户端的请求。
}

func Init() {
	setup.InitConfig()                         // 初始化 config.yaml 配置文件
	config.GormConfig = setup.InitGormConfig() // 初始化 GORM 配置
	config.Database = setup.InitGorm()         // 初始化数据库连接
	if config.Database != nil {
		if err := setup.CreateTable(config.Database); err != nil {
			log.Printf("无法创建或更新表: %v", err)
		}
		log.Println("数据库表初始化成功!")
	} else {
		log.Fatalf("数据库连接为空，无法创建或更新表，请检查...")
	}
}
