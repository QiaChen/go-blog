// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Interact is the golang structure for table interact.
type Interact struct {
	Id         uint        `json:"id"         orm:"id"          description:"自增ID"`
	Type       int         `json:"type"       orm:"type"        description:"操作类型。0:赞，1:踩。"`
	UserId     uint        `json:"userId"     orm:"user_id"     description:"操作用户"`
	TargetId   uint        `json:"targetId"   orm:"target_id"   description:"对应内容ID，该内容可能是content, reply"`
	TargetType string      `json:"targetType" orm:"target_type" description:"内容模型: content, reply, 具体由程序定义"`
	Count      uint        `json:"count"      orm:"count"       description:"操作数据值"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:""`
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:""`
}
