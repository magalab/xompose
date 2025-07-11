// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"context"
	"xompose/internal/dao/internal"
	"xompose/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// userDao is the data access object for the table user.
// You can define custom methods on it to extend its functionality as needed.
type userDao struct {
	*internal.UserDao
}

var (
	// User is a globally accessible object for table user operations.
	User = userDao{internal.NewUserDao()}
)

// Add your custom methods and functionality below.
// GetUserByName 根据用户名获取用户
func (d *userDao) GetUserByName(ctx context.Context, name string) (user *model.UserModel, err error) {
	var item *model.UserModel
	if err := d.Ctx(ctx).
		Where(d.Columns().Username, name).
		Scan(&item); err != nil {
		return nil, err
	}

	return item, nil
}

// GetUserById 根据用户ID获取用户
func (d *userDao) GetUserById(ctx context.Context, id int) (user *model.UserModel, err error) {
	var item *model.UserModel
	if err := d.Ctx(ctx).
		Where(d.Columns().Id, id).
		Scan(&item); err != nil {
		return nil, err
	}

	return item, nil
}

// Count 统计用户数量
func (d *userDao) Count(ctx context.Context) (int, error) {
	return d.Ctx(ctx).Count()
}

// UpdatePasswordById 更新密码
func (d *userDao) UpdatePasswordById(ctx context.Context, id int, password string) error {
	updater := g.Map{
		d.Columns().Password: password,
	}
	if _, err := d.Ctx(ctx).
		Where(d.Columns().Id, id).
		UpdateAndGetAffected(updater); err != nil {
		return err
	}

	return nil
}
