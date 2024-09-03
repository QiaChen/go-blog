// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type (
	ISetting interface {
		// 设置KV。
		Set(ctx context.Context, key string, value string) error
		// 查询KV。
		Get(ctx context.Context, key string) (string, error)
		// 查询KV，返回泛型，便于转换。
		GetVar(ctx context.Context, key string) (*g.Var, error)
	}
)

var (
	localSetting ISetting
)

func Setting() ISetting {
	if localSetting == nil {
		panic("implement not found for interface ISetting, forgot register?")
	}
	return localSetting
}

func RegisterSetting(i ISetting) {
	localSetting = i
}
