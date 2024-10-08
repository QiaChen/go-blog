// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ContentDao is the data access object for table gf_content.
type ContentDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns ContentColumns // columns contains all the column names of Table for convenient usage.
}

// ContentColumns defines and stores column names for table gf_content.
type ContentColumns struct {
	Id             string // 自增ID
	Key            string // 唯一键名，用于程序硬编码，一般不常用
	Type           string // 内容模型: topic, ask, article等，具体由程序定义
	CategoryId     string // 栏目ID
	UserId         string // 用户ID
	AdoptedReplyId string // 采纳的回复ID，问答模块有效
	Title          string // 标题
	Content        string // 内容
	Sort           string // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
	Brief          string // 摘要
	Thumb          string // 缩略图
	Tags           string // 标签名称列表，以JSON存储
	Referer        string // 内容来源，例如github/gitee
	Status         string // 状态 0: 正常, 1: 禁用
	ReplyCount     string // 回复数量
	ViewCount      string // 浏览数量
	ZanCount       string // 赞
	CaiCount       string // 踩
	CreatedAt      string // 创建时间
	UpdatedAt      string // 修改时间
}

// contentColumns holds the columns for table gf_content.
var contentColumns = ContentColumns{
	Id:             "id",
	Key:            "key",
	Type:           "type",
	CategoryId:     "category_id",
	UserId:         "user_id",
	AdoptedReplyId: "adopted_reply_id",
	Title:          "title",
	Content:        "content",
	Sort:           "sort",
	Brief:          "brief",
	Thumb:          "thumb",
	Tags:           "tags",
	Referer:        "referer",
	Status:         "status",
	ReplyCount:     "reply_count",
	ViewCount:      "view_count",
	ZanCount:       "zan_count",
	CaiCount:       "cai_count",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewContentDao creates and returns a new DAO object for table data access.
func NewContentDao() *ContentDao {
	return &ContentDao{
		group:   "default",
		table:   "gf_content",
		columns: contentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ContentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ContentDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ContentDao) Columns() ContentColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ContentDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ContentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ContentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
