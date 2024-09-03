// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Content is the golang structure for table content.
type Content struct {
	Id             uint        `json:"id"             orm:"id"               description:"自增ID"`
	Key            string      `json:"key"            orm:"key"              description:"唯一键名，用于程序硬编码，一般不常用"`
	Type           string      `json:"type"           orm:"type"             description:"内容模型: topic, ask, article等，具体由程序定义"`
	CategoryId     uint        `json:"categoryId"     orm:"category_id"      description:"栏目ID"`
	UserId         uint        `json:"userId"         orm:"user_id"          description:"用户ID"`
	AdoptedReplyId uint        `json:"adoptedReplyId" orm:"adopted_reply_id" description:"采纳的回复ID，问答模块有效"`
	Title          string      `json:"title"          orm:"title"            description:"标题"`
	Content        string      `json:"content"        orm:"content"          description:"内容"`
	Sort           uint        `json:"sort"           orm:"sort"             description:"排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶"`
	Brief          string      `json:"brief"          orm:"brief"            description:"摘要"`
	Thumb          string      `json:"thumb"          orm:"thumb"            description:"缩略图"`
	Tags           string      `json:"tags"           orm:"tags"             description:"标签名称列表，以JSON存储"`
	Referer        string      `json:"referer"        orm:"referer"          description:"内容来源，例如github/gitee"`
	Status         uint        `json:"status"         orm:"status"           description:"状态 0: 正常, 1: 禁用"`
	ReplyCount     uint        `json:"replyCount"     orm:"reply_count"      description:"回复数量"`
	ViewCount      uint        `json:"viewCount"      orm:"view_count"       description:"浏览数量"`
	ZanCount       uint        `json:"zanCount"       orm:"zan_count"        description:"赞"`
	CaiCount       uint        `json:"caiCount"       orm:"cai_count"        description:"踩"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"       description:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"       description:"修改时间"`
}
