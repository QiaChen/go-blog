// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SettingDao is the data access object for table gf_setting.
type SettingDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns SettingColumns // columns contains all the column names of Table for convenient usage.
}

// SettingColumns defines and stores column names for table gf_setting.
type SettingColumns struct {
	K         string // 键名
	V         string // 键值
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// settingColumns holds the columns for table gf_setting.
var settingColumns = SettingColumns{
	K:         "k",
	V:         "v",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSettingDao creates and returns a new DAO object for table data access.
func NewSettingDao() *SettingDao {
	return &SettingDao{
		group:   "default",
		table:   "gf_setting",
		columns: settingColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SettingDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SettingDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SettingDao) Columns() SettingColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SettingDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SettingDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SettingDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
