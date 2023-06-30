package controller

import (
	"first_work_jty/service"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	service.UserRegister(c, username, password)
}

func UserLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	service.UserLogin(c, username, password)
}

func UserUpdatePassword(c *gin.Context) {
	username := c.PostForm("username")
	oldPassword := c.PostForm("oldPassword")
	newPassword := c.PostForm("newPassword")
	service.UserUpdatePassword(c, username, oldPassword, newPassword)
}

func UserDelete(c *gin.Context) {
	id := c.PostForm("id")
	service.UserDelete(c, id)
}
