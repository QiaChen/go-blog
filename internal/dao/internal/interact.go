// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// InteractDao is the data access object for table gf_interact.
type InteractDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns InteractColumns // columns contains all the column names of Table for convenient usage.
}

// InteractColumns defines and stores column names for table gf_interact.
type InteractColumns struct {
	Id         string // 自增ID
	Type       string // 操作类型。0:赞，1:踩。
	UserId     string // 操作用户
	TargetId   string // 对应内容ID，该内容可能是content, reply
	TargetType string // 内容模型: content, reply, 具体由程序定义
	Count      string // 操作数据值
	CreatedAt  string //
	UpdatedAt  string //
}

// interactColumns holds the columns for table gf_interact.
var interactColumns = InteractColumns{
	Id:         "id",
	Type:       "type",
	UserId:     "user_id",
	TargetId:   "target_id",
	TargetType: "target_type",
	Count:      "count",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewInteractDao creates and returns a new DAO object for table data access.
func NewInteractDao() *InteractDao {
	return &InteractDao{
		group:   "default",
		table:   "gf_interact",
		columns: interactColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *InteractDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *InteractDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *InteractDao) Columns() InteractColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *InteractDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *InteractDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *InteractDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
