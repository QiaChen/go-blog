// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Category is the golang structure for table category.
type Category struct {
	Id          uint        `json:"id"          orm:"id"           description:"分类ID，自增主键"`
	ContentType string      `json:"contentType" orm:"content_type" description:"内容类型：topic, ask, article, reply"`
	Key         string      `json:"key"         orm:"key"          description:"栏目唯一键名，用于程序部分场景硬编码，一般不会用得到"`
	ParentId    uint        `json:"parentId"    orm:"parent_id"    description:"父级分类ID，用于层级管理"`
	UserId      uint        `json:"userId"      orm:"user_id"      description:"创建的用户ID"`
	Name        string      `json:"name"        orm:"name"         description:"分类名称"`
	Sort        uint        `json:"sort"        orm:"sort"         description:"排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶"`
	Thumb       string      `json:"thumb"       orm:"thumb"        description:"封面图"`
	Brief       string      `json:"brief"       orm:"brief"        description:"简述"`
	Content     string      `json:"content"     orm:"content"      description:"详细介绍"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"修改时间"`
}
