// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        uint        `json:"id"        orm:"id"         description:"UID"`
	Passport  string      `json:"passport"  orm:"passport"   description:"账号"`
	Password  string      `json:"password"  orm:"password"   description:"MD5密码"`
	Nickname  string      `json:"nickname"  orm:"nickname"   description:"昵称"`
	Avatar    string      `json:"avatar"    orm:"avatar"     description:"头像地址"`
	Status    int         `json:"status"    orm:"status"     description:"状态 0:启用 1:禁用"`
	Gender    int         `json:"gender"    orm:"gender"     description:"性别 0: 未设置 1: 男 2: 女"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"注册时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
}
