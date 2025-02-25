package main

import (
	"backend/config"
	"backend/routes"
	"log"
)

func main() {
	//初始化配置
	config.InitConfig();

	if config.AppConfig == nil {
        log.Fatalf("bull shit!AppConfig is not initialized")
    }

	// 初始化数据库连接
    if err := config.InitDB(); err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }

    // 设置路由
    r := routes.SetupRouter()

    // 启动服务器
    port := config.AppConfig.App.Port
    if port == "" {
        port = "8080" // 默认端口
    }
    if err := r.Run(":" + port); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}