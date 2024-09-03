package i

import (
	"context"
	"go-blog/api/v1"
	"go-blog/internal/model"
	"go-blog/internal/service"
)

// 注册控制器
var Page = cPage{}

type cPage struct{}

func (a *cPage) Index(ctx context.Context, req *v1.PageUserIndexReq) (res *v1.PageUserIndexRes, err error) {
	data, _ := service.Page().List(ctx)
	service.View().RenderTpl(ctx, "index/i/user_layout.html", model.View{
		Data:    data,
		MainTpl: "index/i/page/index.html",
	})
	return
}

func (a *cPage) Create(ctx context.Context, req *v1.PageUserCreateReq) (res *v1.PageUserCreateRes, err error) {
	service.View().RenderTpl(ctx, "index/i/user_layout.html", model.View{
		MainTpl: "index/i/page/create.html",
	})
	return
}

func (a *cPage) CreateDo(ctx context.Context, req *v1.PageUserCreateDoReq) (res *v1.PageUserCreateDoRes, err error) {
	out, err := service.Page().Create(ctx, model.PageCreateInput{
		PageCreateUpdateBase: model.PageCreateUpdateBase{
			Title:       req.Title,
			UrlPath:     req.UrlPath,
			Description: req.Description,
			Content:     req.Content,
			Template:    req.Template,
			Reply:       req.Reply,
		},
	})
	if err != nil {
		return nil, err
	}
	return &v1.PageUserCreateDoRes{Id: out.Id}, nil
	return
}

func (a *cPage) Update(ctx context.Context, req *v1.PageUserUpdateReq) (res *v1.PageUserUpdateRes, err error) {
	data, _ := service.Page().GetPageById(ctx, req.Id)
	service.View().RenderTpl(ctx, "index/i/user_layout.html", model.View{
		Data:    data,
		MainTpl: "index/i/page/update.html",
	})
	return
}

func (a *cPage) UpdateDo(ctx context.Context, req *v1.PageUserUpdateDoReq) (res *v1.PageUserUpdateDoRes, err error) {

	err = service.Page().Update(ctx, model.PageUpdateInput{
		Id: req.Id,
		PageCreateUpdateBase: model.PageCreateUpdateBase{
			Title:       req.Title,
			UrlPath:     req.UrlPath,
			Description: req.Description,
			Content:     req.Content,
			Template:    req.Template,
			Reply:       req.Reply,
		},
	})
	res = &v1.PageUserUpdateDoRes{Id: req.Id}
	return
}

func (A *cPage) Delete(ctx context.Context, req *v1.PageDeleteReq) (res *v1.PagetDeleteRes, err error) {
	err = service.Page().Delete(ctx, req.Id)
	return
}
