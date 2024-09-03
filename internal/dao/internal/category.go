// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CategoryDao is the data access object for table gf_category.
type CategoryDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns CategoryColumns // columns contains all the column names of Table for convenient usage.
}

// CategoryColumns defines and stores column names for table gf_category.
type CategoryColumns struct {
	Id          string // 分类ID，自增主键
	ContentType string // 内容类型：topic, ask, article, reply
	Key         string // 栏目唯一键名，用于程序部分场景硬编码，一般不会用得到
	ParentId    string // 父级分类ID，用于层级管理
	UserId      string // 创建的用户ID
	Name        string // 分类名称
	Sort        string // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
	Thumb       string // 封面图
	Brief       string // 简述
	Content     string // 详细介绍
	CreatedAt   string // 创建时间
	UpdatedAt   string // 修改时间
}

// categoryColumns holds the columns for table gf_category.
var categoryColumns = CategoryColumns{
	Id:          "id",
	ContentType: "content_type",
	Key:         "key",
	ParentId:    "parent_id",
	UserId:      "user_id",
	Name:        "name",
	Sort:        "sort",
	Thumb:       "thumb",
	Brief:       "brief",
	Content:     "content",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewCategoryDao creates and returns a new DAO object for table data access.
func NewCategoryDao() *CategoryDao {
	return &CategoryDao{
		group:   "default",
		table:   "gf_category",
		columns: categoryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CategoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CategoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CategoryDao) Columns() CategoryColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CategoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CategoryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CategoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
