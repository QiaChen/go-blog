// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Reply is the golang structure for table reply.
type Reply struct {
	Id         uint        `json:"id"         orm:"id"          description:"回复ID"`
	ParentId   uint        `json:"parentId"   orm:"parent_id"   description:"回复对应的上一级回复ID(没有的话默认为0)"`
	Title      string      `json:"title"      orm:"title"       description:"回复标题"`
	Content    string      `json:"content"    orm:"content"     description:"回复内容"`
	TargetType string      `json:"targetType" orm:"target_type" description:"评论类型: content, reply"`
	TargetId   uint        `json:"targetId"   orm:"target_id"   description:"对应内容ID，可能回复的是另一个回复，所以这里没有使用content_id"`
	UserId     uint        `json:"userId"     orm:"user_id"     description:"网站用户ID"`
	ZanCount   uint        `json:"zanCount"   orm:"zan_count"   description:"赞"`
	CaiCount   uint        `json:"caiCount"   orm:"cai_count"   description:"踩"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:""`
}
