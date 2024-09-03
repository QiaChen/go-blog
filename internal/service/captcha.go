// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	ICaptcha interface {
		// NewAndStore 创建验证码，直接输出验证码图片内容到HTTP Response.
		NewAndStore(ctx context.Context, name string) error
		// VerifyAndClear 校验验证码，并清空缓存的验证码信息
		VerifyAndClear(r *ghttp.Request, name string, value string) bool
	}
)

var (
	localCaptcha ICaptcha
)

func Captcha() ICaptcha {
	if localCaptcha == nil {
		panic("implement not found for interface ICaptcha, forgot register?")
	}
	return localCaptcha
}

func RegisterCaptcha(i ICaptcha) {
	localCaptcha = i
}
