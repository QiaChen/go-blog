// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Page is the golang structure for table page.
type Page struct {
	Id          int         `json:"id"          orm:"id"          description:"id"`
	Title       string      `json:"title"       orm:"title"       description:"标题"`
	UrlPath     string      `json:"urlPath"     orm:"url_path"    description:"描述"`
	Description string      `json:"description" orm:"description" description:"介绍"`
	Content     string      `json:"content"     orm:"content"     description:"内容"`
	Template    string      `json:"template"    orm:"template"    description:"页面模版"`
	Reply       int         `json:"reply"       orm:"reply"       description:"是否开启评论"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:""`
}
