package service

import (
	"first_work_jty/model"
	"first_work_jty/util/response"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context, username, password string) {
	if username == "" || password == "" {
		response.Failed(c, "用户名，密码不得为空！")
		return
	}
	user := model.GetUserByName(username)
	if user.Username == username {
		response.Failed(c, "用户名已被占用！")
		return
	}
	user = &model.User{
		Username: username,
		Password: password,
	}
	if err := model.CreateUser(user); err != nil {
		response.Error(c, err.Error())
	}
	response.OK(c, user)
}

func UserLogin(c *gin.Context, username, password string) {
	if username == "" || password == "" {
		response.Failed(c, "用户名，密码不得为空！")
		return
	}
	user := model.GetUserByName(username)
	if username != user.Username {
		response.Failed(c, "用户名不存在！")
		return
	}
	if user.Password != password {
		response.Failed(c, "密码错误！")
		return
	}
	response.OK(c, user)
}

func UserUpdatePassword(c *gin.Context, username, oldPassword, newPassword string) {
	if username == "" || oldPassword == "" || newPassword == "" {
		response.Failed(c, "用户名，密码不得为空！")
		return
	}
	user := model.GetUserByName(username)
	if username != user.Username {
		response.Failed(c, "用户名不存在！")
		return
	}
	if user.Password != oldPassword {
		response.Failed(c, "旧密码错误！")
		return
	}
	if user.Password == newPassword {
		response.Failed(c, "新旧密码不能相同！")
		return
	}
	user.Password = newPassword
	if err := model.UpdateUser(user); err != nil {
		response.Error(c, err.Error())
		return
	}
	response.OK(c, nil)
}

func UserDelete(c *gin.Context, id string) {
	if err := model.DeleteUserById(id); err != nil {
		response.Error(c, err.Error())
	}
}
