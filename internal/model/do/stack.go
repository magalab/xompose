// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Stack is the golang structure of table stack for DAO operations like Where/Data.
type Stack struct {
	g.Meta      `orm:"table:stack, do:true"`
	Id          interface{} // 主键 ID
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   interface{} // 软删除
	StackName   interface{} // 服务名称
	YamlContent interface{} // yaml 文件
	StackStatus interface{} // 服务状态
	YamlPath    interface{} // YAML文件路劲
	EnvContent  interface{} // 环境变量
}
