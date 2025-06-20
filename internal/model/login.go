package model

type PhoneLoginReq struct {
	PhoneNumber string `v:"phone-loose"`
	Password    string `v:"required#不能为空"`
}

type UsernameLoginReq struct {
	Username string `v:"required#不能为空"`
	Password string `v:"required#不能为空"`
}

type DoLoginReq struct {
	Id int
}

type RegisterReq struct {
	Username  string `v:"required#不能为空"`
	Password  string `v:"required#不能为空"`
	Password2 string `v:"required#不能为空"`
}
