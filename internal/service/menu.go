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
	IMenu interface {
		// 获取顶部菜单
		SetTopMenus(ctx context.Context, menus []*model.MenuItem) error
		// 获取顶部菜单
		GetTopMenus(ctx context.Context) ([]*model.MenuItem, error)
		// 根据给定的Url检索顶部菜单，给定的Url可能只是一个Url Path。
		GetTopMenuByUrl(ctx context.Context, url string) (*model.MenuItem, error)
	}
)

var (
	localMenu IMenu
)

func Menu() IMenu {
	if localMenu == nil {
		panic("implement not found for interface IMenu, forgot register?")
	}
	return localMenu
}

func RegisterMenu(i IMenu) {
	localMenu = i
}
