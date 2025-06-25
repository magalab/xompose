package v1

import (
	"xompose/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// type NetworkAddReq struct {
//   	g.Meta `path:"/network" tags:"Network" method:"post" sm:"TBD"`
//   	// 数据
// }
// type NetworkAddRes struct {
// 	// 数据
// }

// type NetworkGetReq struct {
//   	g.Meta `path:"/network" tags:"Network" method:"get" sm:"TBD"`
//   	// 数据
// }
// type NetworkGetRes struct {
// 	// 数据
// }

type NetworkListReq struct {
	g.Meta `path:"/networks" tags:"Network" method:"get" sm:"获取所有的 network"`
}
type NetworkListRes struct {
	// 数据
	Items []*model.NetworkItem `json:"items"`
}

// type NetworkUpdateReq struct {
//   	g.Meta `path:"/network" tags:"Network" method:"put" sm:"TBD"`
//   	// 数据
// }
// type NetworkUpdateRes struct {
// 	// 数据
// }

// type NetworkDeleteReq struct {
//   	g.Meta `path:"/network" tags:"Network" method:"delete" sm:"TBD"`
//   	// 数据
// }
// type NetworkDeleteRes struct {
// 	// 数据
// }
