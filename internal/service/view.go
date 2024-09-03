// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"go-blog/internal/model"
)

type (
	IView interface {
		// GetBreadCrumb 前台系统-获取面包屑列表.
		GetBreadCrumb(ctx context.Context, in *model.ViewGetBreadCrumbInput) []model.ViewBreadCrumb
		// GetTitle 前台系统-获取标题
		GetTitle(ctx context.Context, in *model.ViewGetTitleInput) string
		// RenderTpl 渲染指定模板页面
		RenderTpl(ctx context.Context, tpl string, data ...model.View)
		// Render 渲染默认模板页面
		Render(ctx context.Context, data ...model.View)
		// Render302 跳转中间页面
		Render302(ctx context.Context, data ...model.View)
		// Render401 401页面
		Render401(ctx context.Context, data ...model.View)
		// Render403 403页面
		Render403(ctx context.Context, data ...model.View)
		// Render404 404页面
		Render404(ctx context.Context, data ...model.View)
		// Render500 500页面
		Render500(ctx context.Context, data ...model.View)
	}
)

var (
	localView IView
)

func View() IView {
	if localView == nil {
		panic("implement not found for interface IView, forgot register?")
	}
	return localView
}

func RegisterView(i IView) {
	localView = i
}
