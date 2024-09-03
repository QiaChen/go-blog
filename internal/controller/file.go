package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"go-blog/api/v1"
	"go-blog/internal/model"
	"go-blog/internal/service"
)

// 文件管理
var File = cFile{}

type cFile struct{}

func (a *cFile) Upload(ctx context.Context, req *v1.FileUploadReq) (res *v1.FileUploadRes, err error) {
	if req.File == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请选择需要上传的文件")
	}
	result, err := service.File().Upload(ctx, model.FileUploadInput{
		File:       req.File,
		RandomName: true,
	})
	if err != nil {
		return nil, err
	}

	res = &v1.FileUploadRes{
		Name: result.Name,
		Url:  result.Url,
	}
	r := g.RequestFromCtx(ctx)
	r.Response.Write("{\"success\":1, \"message\":\"ok\", \"url\": \" " + result.Url + "\" }")
	r.Exit()

	return
}
