// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Page is the golang structure of table gf_page for DAO operations like Where/Data.
type Page struct {
	g.Meta      `orm:"table:gf_page, do:true"`
	Id          interface{} // id
	Title       interface{} // 标题
	UrlPath     interface{} // 描述
	Description interface{} // 介绍
	Content     interface{} // 内容
	Template    interface{} // 页面模版
	Reply       interface{} // 是否开启评论
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time //
}
