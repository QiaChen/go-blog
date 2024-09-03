// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"go-blog/internal/model"
	"go-blog/internal/model/entity"
)

type (
	IPage interface {
		// 通过id获取页面信息
		GetPageById(ctx context.Context, id uint) (out *entity.Page, err error)
		// 通过path获取页面信息
		GetPageByPath(ctx context.Context, path string) (out *entity.Page, err error)
		// Create 创建内容
		Create(ctx context.Context, in model.PageCreateInput) (out model.PageCreateOutput, err error)
		List(ctx context.Context) (out model.PagetGetListOutput, err error)
		Update(ctx context.Context, in model.PageUpdateInput) error
		Delete(ctx context.Context, id uint) error
	}
)

var (
	localPage IPage
)

func Page() IPage {
	if localPage == nil {
		panic("implement not found for interface IPage, forgot register?")
	}
	return localPage
}

func RegisterPage(i IPage) {
	localPage = i
}
