package blog

import (
	"context"

	"go-blog/api/v1"
	"go-blog/internal/consts"
	"go-blog/internal/model"
	"go-blog/internal/service"
)

// Article 文章管理
var Blog = cBlog{}

type cBlog struct{}

// Index article list
func (a *cBlog) Index(ctx context.Context, req *v1.BlogIndexReq) (res *v1.BlogIndexRes, err error) {
	req.Type = consts.ContentTypeBlog
	getListRes, err := service.Content().GetList(ctx, model.ContentGetListInput{
		Type:       req.Type,
		CategoryId: req.CategoryId,
		Page:       req.Page,
		Size:       req.Size,
		Sort:       req.Sort,
	})
	if err != nil {
		return nil, err
	}
	service.View().Render(ctx, model.View{
		MainTpl:     "index/blog/index2.html",
		ContentType: req.Type,
		Data:        getListRes,
		Title: service.View().GetTitle(ctx, &model.ViewGetTitleInput{
			ContentType: req.Type,
			CategoryId:  req.CategoryId,
		}),
	})
	return
}

// Detail .article details
func (a *cBlog) Detail(ctx context.Context, req *v1.BlogDetailReq) (res *v1.BlogDetailRes, err error) {
	getDetailRes, err := service.Content().GetDetail(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if getDetailRes == nil {
		service.View().Render404(ctx)
		return nil, nil
	}
	if err = service.Content().AddViewCount(ctx, req.Id, 1); err != nil {
		return res, err
	}
	service.View().Render(ctx, model.View{
		MainTpl:     "index/blog/detail.html",
		ContentType: consts.ContentTypeBlog,
		Data:        getDetailRes,
		Title: service.View().GetTitle(ctx, &model.ViewGetTitleInput{
			ContentType: getDetailRes.Content.Type,
			CategoryId:  getDetailRes.Content.CategoryId,
			CurrentName: getDetailRes.Content.Title,
		}),
		BreadCrumb: service.View().GetBreadCrumb(ctx, &model.ViewGetBreadCrumbInput{
			ContentId:   getDetailRes.Content.Id,
			ContentType: getDetailRes.Content.Type,
			CategoryId:  getDetailRes.Content.CategoryId,
		}),
	})
	return
}
