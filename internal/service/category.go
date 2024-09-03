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
	ICategory interface {
		// GetTree 查询列表
		GetTree(ctx context.Context, contentType string) ([]*model.CategoryTreeItem, error)
		// GetSubIdList 获取指定栏目ID及其下面所有子ID，构成数组返回。
		// 注意，返回的ID列表中包含查询的栏目ID.
		GetSubIdList(ctx context.Context, id uint) ([]uint, error)
		// GetList 获得所有的栏目列表。
		GetList(ctx context.Context) (list []*entity.Category, err error)
		// GetItem 查询单个栏目信息
		GetItem(ctx context.Context, id uint) (*entity.Category, error)
		// GetMap 获得所有的栏目列表，构成Map返回，键名为栏目ID
		GetMap(ctx context.Context) (map[uint]*entity.Category, error)
	}
)

var (
	localCategory ICategory
)

func Category() ICategory {
	if localCategory == nil {
		panic("implement not found for interface ICategory, forgot register?")
	}
	return localCategory
}

func RegisterCategory(i ICategory) {
	localCategory = i
}
