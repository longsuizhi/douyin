package main

import (
	"douyin/conf"
	"douyin/dao"
	"douyin/logger"
	"douyin/router"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	config := "app/config/config.yaml"
	if len(os.Args) > 2 {
		fmt.Println("need config path, eg.: ./douyin conf/config.yaml")
		config = os.Args[1]
		return
	}

	// 1. 加载配置
	if err := conf.Init(config); err != nil {
		fmt.Printf("init settings failed, err = %v\n", err)
		return
	}
	// 2. 初始化日志
	if err := logger.Init(conf.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err = %v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success")

	// 3. 初始化 MySQL 连接
	dao.InitDB()
	defer dao.MysqlClose()

	// 4. 初始化 Redis 连接
	if err := dao.InitRedisClient(); err != nil {
		fmt.Printf("init logger failed, err = %v\n", err)
		return
	}
	defer dao.RedisClose()

	// 5.获取Engine
	c := gin.Default()

	// 6.注册路由
	router.InitRouter(c)

	// 7.运行服务
	c.Run(fmt.Sprintf(":%d", conf.Info.Port))
}
