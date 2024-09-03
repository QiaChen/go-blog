// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// File is the golang structure for table file.
type File struct {
	Id        uint        `json:"id"        orm:"id"         description:"自增ID"`
	Name      string      `json:"name"      orm:"name"       description:"文件名称"`
	Src       string      `json:"src"       orm:"src"        description:"本地文件存储路径"`
	Url       string      `json:"url"       orm:"url"        description:"URL地址，可能为空"`
	UserId    uint        `json:"userId"    orm:"user_id"    description:"操作用户"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
}
