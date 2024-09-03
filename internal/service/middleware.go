// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IMiddleware interface {
		// 返回处理中间件
		ResponseHandler(r *ghttp.Request)
		// 自定义上下文对象
		Ctx(r *ghttp.Request)
		// 前台系统权限控制，用户必须登录才能访问
		Auth(r *ghttp.Request)
		// 只有管理员才能访问
		AuthAdmin(r *ghttp.Request)
	}
)

var (
	localMiddleware IMiddleware
)

func Middleware() IMiddleware {
	if localMiddleware == nil {
		panic("implement not found for interface IMiddleware, forgot register?")
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}
