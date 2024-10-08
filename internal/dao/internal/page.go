// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PageDao is the data access object for table gf_page.
type PageDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns PageColumns // columns contains all the column names of Table for convenient usage.
}

// PageColumns defines and stores column names for table gf_page.
type PageColumns struct {
	Id          string // id
	Title       string // 标题
	UrlPath     string // 描述
	Description string // 介绍
	Content     string // 内容
	Template    string // 页面模版
	Reply       string // 是否开启评论
	CreatedAt   string // 创建时间
	UpdatedAt   string //
}

// pageColumns holds the columns for table gf_page.
var pageColumns = PageColumns{
	Id:          "id",
	Title:       "title",
	UrlPath:     "url_path",
	Description: "description",
	Content:     "content",
	Template:    "template",
	Reply:       "reply",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewPageDao creates and returns a new DAO object for table data access.
func NewPageDao() *PageDao {
	return &PageDao{
		group:   "default",
		table:   "gf_page",
		columns: pageColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PageDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PageDao) Columns() PageColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PageDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PageDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
