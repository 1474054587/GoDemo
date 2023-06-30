package main

import (
	"first_work_jty/config"
	"first_work_jty/dao"
	"first_work_jty/model"
	"first_work_jty/router"
	"fmt"
)

func main() {
	// 加载配置文件
	if err := config.Init("./config/config.ini"); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}

	// 连接数据库
	if err := dao.InitMySQL(config.Config.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	// 创建表
	dao.DB.AutoMigrate(&model.User{})
	// 断开数据库连接
	defer dao.Close()

	// 注册路由
	router.InitRouter()

	// 启动服务
	if err := router.Run(); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}

}
