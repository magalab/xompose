package model

import "xompose/internal/model/entity"

type UserModel struct {
	entity.User
}

type UserItem struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func (m *UserItem) IsLogin() bool {
	return m.Id >= 0
}

func (m *UserModel) ToItem() *UserItem {
	return &UserItem{
		Id:       m.Id,
		Username: m.Username,
	}
}

// type UserAddReq struct {
// 	PhoneNumber string `json:"phoneNumber" v:"required|phone-loose"`
// 	Username    string `json:"username" v:"required#姓名不能为空"`
// 	Password    string `json:"password" v:"required#密码不能为空"`
// }

// type UserGetReq struct {
// 	Id int `json:"userId" v:"required"`
// }

// type RegisterReq struct {
// 	PhoneNumber string `json:"phoneNumber" v:"required"`
// }

type ChangePasswordReq struct {
	CurrentPassword string `v:"required#请输入旧密码" json:"currentPassword" dc:"旧密码"`
	NewPassword     string `v:"password2#密码不符合要求" json:"newPassword" dc:"新密码"`
	NewPassword2    string `v:"password2|same:NewPassword#二次密码不一致" json:"newPassword2" dc:"确认密码"`
	Id              int    `json:"-"`
}
