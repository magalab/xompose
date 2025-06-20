// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Stack is the golang structure for table stack.
type Stack struct {
	Id          int         `json:"id"          orm:"id"           ` // 主键 ID
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   ` // 更新时间
	DeletedAt   uint64      `json:"deletedAt"   orm:"deleted_at"   ` // 软删除
	StackName   string      `json:"stackName"   orm:"stack_name"   ` // 服务名称
	YamlContent string      `json:"yamlContent" orm:"yaml_content" ` // yaml 文件
	StackStatus string      `json:"stackStatus" orm:"stack_status" ` // 服务状态
	YamlPath    string      `json:"yamlPath"    orm:"yaml_path"    ` // YAML文件路劲
}
