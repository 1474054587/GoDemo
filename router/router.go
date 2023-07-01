package router

import (
	"first_work_jty/config"
	"first_work_jty/controller"
	"first_work_jty/handler"
	"fmt"
	"github.com/gin-gonic/gin"
)

var Engine *gin.Engine

func InitRouter() *gin.Engine {
	if config.Config.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	Engine = gin.Default()

	// 引入静态文件
	Engine.Static("/static", "static")
	// 引入模板文件
	Engine.LoadHTMLGlob("templates/*")

	Engine.GET("/", handler.LoginIndexHandler)

	loginGroup := Engine.Group("/login")
	{
		loginGroup.GET("/", handler.LoginIndexHandler)
		loginGroup.POST("/", controller.UserLogin)
		loginGroup.POST("/register", controller.UserRegister)
		loginGroup.POST("/updatePassword", controller.UserUpdatePassword)
		loginGroup.POST("/deleteUser", controller.UserDelete)
	}
	return Engine
}

func Run() error {
	return Engine.Run(fmt.Sprintf(":%d", config.Config.Port))
}
