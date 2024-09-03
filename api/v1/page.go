package v1

import "github.com/gogf/gf/v2/frame/g"

type PageUserIndexReq struct {
	g.Meta `path:"/page" method:"get" summary:"页面管理" tags:"页面"`
}
type PageUserIndexRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type PageUserCreateReq struct {
	g.Meta `path:"/page/create" method:"get" summary:"添加页面" tags:"页面"`
}
type PageUserCreateRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type PageUserCreateDoReq struct {
	g.Meta      `path:"/page/create" method:"post" summary:"添加页面" tags:"页面"`
	Title       string `json:"title" v:"required#请输入title"   dc:"标题"`
	UrlPath     string `json:"url_path" v:"required#请输入url_path"   dc:"url_path"`
	Description string `json:"description"  dc:"description" d:""`
	Template    string `json:"template" v:"required#请输入template"   dc:"template"`
	Reply       int    `json:"reply"  dc:"reply" d:1`
	Content     string `json:"reply"  dc:"reply" d:""`
}
type PageUserCreateDoRes struct {
	Id uint `v:"min:1#pageid" dc:"pageid"`
}

type PageUserUpdateReq struct {
	g.Meta `path:"/page/update/{Id}" method:"get" tags:"内容" summary:"展示内容修改页面"`
	Id     uint `json:"id" dc:"内容id" v:"min:1#请选择需要修改的内容"`
}
type PageUserUpdateRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type PageUserUpdateDoReq struct {
	g.Meta      `path:"/page/update" method:"post" summary:"编辑页面" tags:"页面"`
	Id          uint   `json:"id" v:"required#请输入id"   dc:"id"`
	Title       string `json:"title" v:"required#请输入title"   dc:"标题"`
	UrlPath     string `json:"url_path" v:"required#请输入url_path"   dc:"url_path"`
	Description string `json:"description"  dc:"description" d:""`
	Template    string `json:"template" v:"required#请输入template"   dc:"template"`
	Reply       int    `json:"reply"  dc:"reply" d:1`
	Content     string `json:"content"  dc:"reply" d:""`
}
type PageUserUpdateDoRes struct {
	Id uint `v:"min:1#pageid" dc:"pageid"`
}
type PageDeleteReq struct {
	g.Meta `path:"/page/delete" method:"delete" tags:"内容" summary:"删除内容接口"`
	Id     uint `v:"min:1#请选择需要删除的内容" dc:"内容id"`
}
type PagetDeleteRes struct{}

type PageShowReq struct {
	g.Meta `path:"/{UrlPath}" method:"get" summary:"页面" tags:"页面"`
}

type PageShowRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
