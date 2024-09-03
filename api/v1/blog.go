package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type BlogIndexReq struct {
	g.Meta `path:"/" method:"get" tags:"文章" summary:"展示Article列表页面"`
	ContentGetListCommonReq
}
type BlogIndexRes struct {
	ContentGetListCommonRes
}

type BlogDetailReq struct {
	g.Meta `path:"/{Id}" method:"get" tags:"文章" summary:"展示Article详情页面" `
	Id     uint `in:"path" v:"min:1#请选择查看的内容" dc:"内容id"`
}
type BlogDetailRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
