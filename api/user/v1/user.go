package v1

import (
	"xompose/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// type UserAddReq struct {
// 	g.Meta `path:"user" tags:"User" method:"post" sm:"TBD"`
// 	// 数据
// }
// type UserAddRes struct {
// 	// 数据
// }

type UserInfoReq struct {
	g.Meta `path:"/user/info" tags:"用户(User)" method:"get" sm:"用户信息"`
}
type UserInfoRes struct {
	*model.UserItem
}

// type UserListReq struct {
//   	g.Meta `path:"users" tags:"User" method:"get" sm:"TBD"`
//   	// 数据
// }
// type UserListRes struct {
// 	// 数据
// }

// type UserUpdateReq struct {
//   	g.Meta `path:"user" tags:"User" method:"put" sm:"TBD"`
//   	// 数据
// }
// type UserUpdateRes struct {
// 	// 数据
// }

// type UserDeleteReq struct {
//   	g.Meta `path:"user" tags:"User" method:"delete" sm:"TBD"`
//   	// 数据
// }
// type UserDeleteRes struct {
// 	// 数据
// }
