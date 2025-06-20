// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        int         `json:"id"        orm:"id"         ` // 主键 ID
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
	DeletedAt uint64      `json:"deletedAt" orm:"deleted_at" ` // 软删除
	Username  string      `json:"username"  orm:"username"   ` // 用户名
	Password  string      `json:"password"  orm:"password"   ` // 密码
}
