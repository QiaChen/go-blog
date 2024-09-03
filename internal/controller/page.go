package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	v1 "go-blog/api/v1"
	"go-blog/internal/model"
	"go-blog/internal/service"
)

var Page = cPage{}

type cPage struct{}

func (a *cPage) Index(ctx context.Context, req *v1.PageShowReq) (res *v1.PageShowRes, err error) {
	r := g.RequestFromCtx(ctx)
	path := r.Get("UrlPath")
	data, err := service.Page().GetPageByPath(ctx, path.String())
	if data == nil {
		r.Response.WriteHeader(404)
		r.Exit()
	}
	service.View().Render(ctx, model.View{
		MainTpl: "index/page/" + data.Template + ".html",
		Data:    data,
		Title: service.View().GetTitle(ctx, &model.ViewGetTitleInput{
			CurrentName: data.Title,
		}),
	})
	return
}
