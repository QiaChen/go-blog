// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Content is the golang structure of table gf_content for DAO operations like Where/Data.
type Content struct {
	g.Meta         `orm:"table:gf_content, do:true"`
	Id             interface{} // 自增ID
	Key            interface{} // 唯一键名，用于程序硬编码，一般不常用
	Type           interface{} // 内容模型: topic, ask, article等，具体由程序定义
	CategoryId     interface{} // 栏目ID
	UserId         interface{} // 用户ID
	AdoptedReplyId interface{} // 采纳的回复ID，问答模块有效
	Title          interface{} // 标题
	Content        interface{} // 内容
	Sort           interface{} // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
	Brief          interface{} // 摘要
	Thumb          interface{} // 缩略图
	Tags           interface{} // 标签名称列表，以JSON存储
	Referer        interface{} // 内容来源，例如github/gitee
	Status         interface{} // 状态 0: 正常, 1: 禁用
	ReplyCount     interface{} // 回复数量
	ViewCount      interface{} // 浏览数量
	ZanCount       interface{} // 赞
	CaiCount       interface{} // 踩
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 修改时间
}
