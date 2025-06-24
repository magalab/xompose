package model

type SystemSetupReq struct {
	Username  string `v:"required#请输入用户名" json:"username" dc:"用户名"`
	Password  string `v:"password2#密码不符合要求" json:"password" dc:"密码"`
	Password2 string `v:"password2|same:Password#二次密码不一致" json:"password2" dc:"确认密码"`
}
