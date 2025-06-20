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
