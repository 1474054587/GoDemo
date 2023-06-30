package model

import (
	"first_work_jty/dao"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"username"`
	Password string `gorm:"password"`
}

func CreateUser(user *User) (err error) {
	err = dao.DB.Create(&user).Error
	return
}

func GetAllUser() (userList []*User, err error) {
	if err = dao.DB.Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

func GetUserById(id string) (user *User, err error) {
	user = new(User)
	err = dao.DB.Debug().Where("id=?", id).First(user).Error
	return
}

func GetUserByName(username string) (user *User) {
	user = new(User)
	dao.DB.Where("username=?", username).First(user)
	return
}

func UpdateUser(user *User) (err error) {
	err = dao.DB.Save(user).Error
	return
}

func DeleteUserById(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&User{}).Error
	return
}
