package VO

type UserVO struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
