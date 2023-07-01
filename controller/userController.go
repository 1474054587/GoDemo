package controller

import (
	"first_work_jty/model/VO"
	"first_work_jty/service"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var userVO VO.UserVO
	c.BindJSON(&userVO)
	username := userVO.Username
	password := userVO.Password
	service.UserRegister(c, username, password)
}

func UserLogin(c *gin.Context) {
	var userVO VO.UserVO
	c.BindJSON(&userVO)
	username := userVO.Username
	password := userVO.Password
	service.UserLogin(c, username, password)
}

func UserUpdatePassword(c *gin.Context) {
	var userVO VO.UserVO
	c.BindJSON(&userVO)
	username := userVO.Username
	oldPassword := userVO.OldPassword
	newPassword := userVO.NewPassword
	service.UserUpdatePassword(c, username, oldPassword, newPassword)
}

func UserDelete(c *gin.Context) {
	var userVO VO.UserVO
	c.BindJSON(&userVO)
	username := userVO.Username
	password := userVO.Password
	service.UserDelete(c, username, password)
}
