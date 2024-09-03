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
	IContent interface {
		// GetList 查询内容列表
		GetList(ctx context.Context, in model.ContentGetListInput) (out *model.ContentGetListOutput, err error)
		// Search 搜索内容列表
		Search(ctx context.Context, in model.ContentSearchInput) (out *model.ContentSearchOutput, err error)
		// GetDetail 查询详情
		GetDetail(ctx context.Context, id uint) (out *model.ContentGetDetailOutput, err error)
		// Create 创建内容
		Create(ctx context.Context, in model.ContentCreateInput) (out model.ContentCreateOutput, err error)
		// Update 修改
		Update(ctx context.Context, in model.ContentUpdateInput) error
		// Delete 删除
		Delete(ctx context.Context, id uint) error
		// AddViewCount 浏览次数增加
		AddViewCount(ctx context.Context, id uint, count int) error
		// AddReplyCount 回复次数增加
		AddReplyCount(ctx context.Context, id uint, count int) error
		// AdoptReply 采纳回复
		AdoptReply(ctx context.Context, id uint, replyID uint) error
		// UnacceptedReply 取消采纳回复
		UnacceptedReply(ctx context.Context, id uint) error
	}
)

var (
	localContent IContent
)

func Content() IContent {
	if localContent == nil {
		panic("implement not found for interface IContent, forgot register?")
	}
	return localContent
}

func RegisterContent(i IContent) {
	localContent = i
}
