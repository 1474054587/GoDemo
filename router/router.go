package router

import (
	"first_work_jty/config"
	"first_work_jty/controller"
	"fmt"
	"github.com/gin-gonic/gin"
)

var Engine *gin.Engine

func InitRouter() *gin.Engine {
	if config.Config.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	Engine = gin.Default()

	loginGroup := Engine.Group("/login")
	{
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
