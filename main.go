package main

import (
	"BAT-douyin/dao/database"
	"BAT-douyin/dao/redis"
	"BAT-douyin/middlewares/logger"
	"BAT-douyin/mypprof"
	"BAT-douyin/routes"
	"BAT-douyin/setting"
	"fmt"
	"go.uber.org/zap"
	"os"
)

func Init() error {
	zap.L().Debug("start loading...")

	//加载配置文件
	if err := setting.Init(os.Args[1]); err != nil {
		zap.L().Error("reading config file failed and err:", zap.Error(err))
		return err
	}
	zap.L().Debug("load configured file success")

	//初始化zap日志
	if err := logger.Init(setting.Conf.LoggerConfig, setting.Conf.AppConfig.Mode); err != nil {
		zap.L().Error("init logger failed...", zap.Error(err))
		return err
	}

	defer zap.L().Sync() // 将缓存中的日志同步到日志文件中
	zap.L().Debug("init logger success!")

	//初始化mysql
	if err := database.Init(setting.Conf.MySQLConfig); err != nil {
		zap.L().Error("init mysql failed...", zap.Error(err))
		return err
	}
	zap.L().Debug("init mysql success")

	//初始化redis
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		zap.L().Error("init redis failed...", zap.Error(err))
		return err
	}
	zap.L().Debug("init redis success")

	//初始化pprof,pprof不是必须的，出错可以正常工作
	go func() {
		err := mypprof.Init(setting.Conf.PprofConfig)
		if err != nil {
			zap.L().Error("pprof start failed..", zap.Error(err))
		}
	}()

	zap.L().Debug("Finish loading configuration file")
	return nil
}

func main() {

	//必须指明yaml配置文件的路径
	if len(os.Args) < 2 {
		fmt.Println("Requires a configuration file path")
		os.Exit(1)
	}

	//启动配置工作
	if err := Init(); err != nil {
		fmt.Println("Failed to load configuration file...")
		os.Exit(1)
	}

	// 注册并启动路由
	if err := routes.Run(setting.Conf.AppConfig); err != nil {
		zap.L().Error("start route failed...", zap.Error(err))
		os.Exit(1)
	}
}
