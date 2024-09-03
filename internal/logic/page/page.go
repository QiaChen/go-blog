package page

import (
	"context"
	"go-blog/internal/dao"
	"go-blog/internal/model"
	"go-blog/internal/model/entity"
	"go-blog/internal/service"
)

type sPage struct{}

func New() *sPage {
	return &sPage{}
}

func init() {
	service.RegisterPage(New())
}

// 通过id获取页面信息
func (s *sPage) GetPageById(ctx context.Context, id uint) (out *entity.Page, err error) {
	err = dao.Page.Ctx(ctx).Where("id", id).Scan(&out)
	return
}

// 通过path获取页面信息
func (s *sPage) GetPageByPath(ctx context.Context, path string) (out *entity.Page, err error) {
	err = dao.Page.Ctx(ctx).Where(dao.Page.Columns().UrlPath, path).Scan(&out)
	return
}

// Create 创建内容
func (s *sPage) Create(ctx context.Context, in model.PageCreateInput) (out model.PageCreateOutput, err error) {

	lastInsertID, err := dao.Page.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.PageCreateOutput{Id: uint(lastInsertID)}, err
}

func (s *sPage) List(ctx context.Context) (out model.PagetGetListOutput, err error) {
	err = dao.Page.Ctx(ctx).Scan(&out.List)
	return
}

func (s *sPage) Update(ctx context.Context, in model.PageUpdateInput) error {
	_, err := dao.Page.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.Content.Columns().Id).
		Where(dao.Content.Columns().Id, in.Id).
		Update()
	return err
}
func (s *sPage) Delete(ctx context.Context, id uint) error {
	_, err := dao.Page.
		Ctx(ctx).
		Delete(dao.Page.Columns().Id, id)
	return err
}
