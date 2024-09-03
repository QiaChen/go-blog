package i

import (
	"context"
	"go-blog/api/v1"
	"go-blog/internal/consts"
	"go-blog/internal/model"
	"go-blog/internal/service"
)

// Content 内容管理
var Content = cContent{}

type cContent struct{}

func (a *cContent) Index(ctx context.Context, req *v1.ContentIndexReq) (res *v1.ContentIndexRes, err error) {
	req.Type = consts.ContentTypeBlog
	err = a.getContentList(ctx, req.UserId, req.ContentGetListCommonReq)
	return
}

func (a *cContent) ShowCreate(ctx context.Context, req *v1.ContentShowCreateReq) (res *v1.ContentShowCreateRes, err error) {
	service.View().RenderTpl(ctx, "index/i/user_layout.html", model.View{
		MainTpl:     "index/i/content/create.html",
		ContentType: req.Type,
	})
	return
}

func (a *cContent) Create(ctx context.Context, req *v1.ContentCreateReq) (res *v1.ContentCreateRes, err error) {
	out, err := service.Content().Create(ctx, model.ContentCreateInput{
		ContentCreateUpdateBase: model.ContentCreateUpdateBase{
			Type:       req.Type,
			CategoryId: req.CategoryId,
			Title:      req.Title,
			Content:    req.Content,
			Brief:      req.Brief,
			Thumb:      req.Thumb,
			Tags:       req.Tags,
			Referer:    req.Referer,
		},
		UserId: service.Session().GetUser(ctx).Id,
	})
	if err != nil {
		return nil, err
	}
	return &v1.ContentCreateRes{ContentId: out.ContentId}, nil
}

func (a *cContent) ShowUpdate(ctx context.Context, req *v1.ContentShowUpdateReq) (res *v1.ContentShowUpdateRes, err error) {
	getDetailRes, err := service.Content().GetDetail(ctx, req.Id)

	if err != nil {
		return nil, err
	}
	service.View().RenderTpl(ctx, "index/i/user_layout.html", model.View{
		MainTpl:     "index/i/content/update.html",
		ContentType: getDetailRes.Content.Type,
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

func (a *cContent) Update(ctx context.Context, req *v1.ContentUpdateReq) (res *v1.ContentUpdateRes, err error) {
	err = service.Content().Update(ctx, model.ContentUpdateInput{
		Id: req.Id,
		ContentCreateUpdateBase: model.ContentCreateUpdateBase{
			Type:       req.Type,
			CategoryId: req.CategoryId,
			Title:      req.Title,
			Content:    req.Content,
			Brief:      req.Brief,
			Thumb:      req.Thumb,
			Tags:       req.Tags,
			Referer:    req.Referer,
		},
	})
	return
}

func (a *cContent) Delete(ctx context.Context, req *v1.ContentDeleteReq) (res *v1.ContentDeleteRes, err error) {
	err = service.Content().Delete(ctx, req.Id)
	return
}

func (a *cContent) getContentList(ctx context.Context, userId uint, req v1.ContentGetListCommonReq) (err error) {
	type getContentListInfo struct {
		Content *model.ContentGetListOutput `json:"content"` // 查询用户
		User    *model.UserGetProfileOutput `json:"i"`       // 查询用户
		Stats   map[string]int              // 发布内容数量
	}
	var (
		data    = getContentListInfo{}
		ctxUser = service.BizCtx().Get(ctx).User
	)
	// 用户内容信息
	data.Content, err = service.Content().GetList(ctx, model.ContentGetListInput{
		Type:       req.Type,
		CategoryId: req.CategoryId,
		Page:       req.Page,
		Size:       req.Size,
		Sort:       req.Sort,
		UserId:     userId,
	})
	if err != nil {
		return err
	}
	// 用户资料信息
	data.User, err = service.User().GetProfileById(ctx, ctxUser.Id)
	if err != nil {
		return err
	}
	// 用户统计信息
	data.Stats, err = service.User().GetUserStats(ctx, ctxUser.Id)
	if err != nil {
		return err
	}

	title := service.View().GetTitle(ctx, &model.ViewGetTitleInput{
		ContentType: req.Type,
		CategoryId:  req.CategoryId,
	})
	service.View().RenderTpl(ctx, "index/i/user_layout.html", model.View{
		MainTpl:     "index/i/content/index.html",
		ContentType: req.Type,
		Data:        data,
		Title:       title,
	})
	return nil
}
