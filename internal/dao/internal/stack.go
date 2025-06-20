// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// StackDao is the data access object for the table stack.
type StackDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  StackColumns       // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// StackColumns defines and stores column names for the table stack.
type StackColumns struct {
	Id          string // 主键 ID
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	DeletedAt   string // 软删除
	StackName   string // 服务名称
	YamlContent string // yaml 文件
	StackStatus string // 服务状态
	YamlPath    string // YAML文件路劲
}

// stackColumns holds the columns for the table stack.
var stackColumns = StackColumns{
	Id:          "id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
	StackName:   "stack_name",
	YamlContent: "yaml_content",
	StackStatus: "stack_status",
	YamlPath:    "yaml_path",
}

// NewStackDao creates and returns a new DAO object for table data access.
func NewStackDao(handlers ...gdb.ModelHandler) *StackDao {
	return &StackDao{
		group:    "default",
		table:    "stack",
		columns:  stackColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *StackDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *StackDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *StackDao) Columns() StackColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *StackDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *StackDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *StackDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
