package service

import (
	"first_work_jty/model/PO"
	"first_work_jty/util/response"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context, username, password string) {
	if username == "" || password == "" {
		response.Failed(c, "用户名，密码不得为空！")
		return
	}
	user := PO.GetUserByName(username)
	if user.Username == username {
		response.Failed(c, "用户名已被占用！")
		return
	}
	user = &PO.User{
		Username: username,
		Password: password,
	}
	if err := PO.CreateUser(user); err != nil {
		response.Error(c, err.Error())
		return
	}
	response.OK(c, user)
}

func UserLogin(c *gin.Context, username, password string) {
	user, ok := compareUsernameAndPassword(c, username, password)
	if !ok {
		return
	}
	response.OK(c, user)
}

func UserUpdatePassword(c *gin.Context, username, oldPassword, newPassword string) {
	user, ok := compareUsernameAndPassword(c, username, oldPassword)
	if !ok {
		return
	}
	if user.Password == newPassword {
		response.Failed(c, "新旧密码不能相同！")
		return
	}
	user.Password = newPassword
	if err := PO.UpdateUser(user); err != nil {
		response.Error(c, err.Error())
		return
	}
	response.OK(c, user)
}

func UserDelete(c *gin.Context, username, password string) {
	user, ok := compareUsernameAndPassword(c, username, password)
	if !ok {
		return
	}
	if err := PO.DeleteUserById(user.ID); err != nil {
		response.Error(c, err.Error())
		return
	}
	response.OK(c, nil)
}

func compareUsernameAndPassword(c *gin.Context, username, password string) (user *PO.User, ok bool) {
	if username == "" || password == "" {
		response.Failed(c, "用户名，密码不得为空！")
		return nil, false
	}
	user = PO.GetUserByName(username)
	if username != user.Username {
		response.Failed(c, "用户名不存在！")
		return nil, false
	}
	if user.Password != password {
		response.Failed(c, "密码错误！")
		return nil, false
	}
	return user, true
}
