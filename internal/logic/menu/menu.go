package menu

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"regexp"

	"go-blog/internal/model"
	"go-blog/internal/service"
)

type sMenu struct{}

const settingTopMenusKey = "TopMenus"

func init() {
	service.RegisterMenu(New())
}

func New() *sMenu {
	return &sMenu{}
}

// 获取顶部菜单
func (s *sMenu) SetTopMenus(ctx context.Context, menus []*model.MenuItem) error {
	b, err := json.Marshal(menus)
	if err != nil {
		return err
	}
	return service.Setting().Set(ctx, settingTopMenusKey, string(b))
}

// 获取顶部菜单
func (s *sMenu) GetTopMenus(ctx context.Context) ([]*model.MenuItem, error) {
	var topMenus []*model.MenuItem
	v, err := service.Setting().GetVar(ctx, settingTopMenusKey)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`\[\[(.*?)\]\]`)

	result := re.ReplaceAllStringFunc(v.String(), func(s string) string {
		app := s[2 : len(s)-2]
		domain := g.Cfg().MustGet(ctx, "setting.domain").String()
		if app == "main" {
			return "//" + domain
		}
		return "//" + app + "." + domain
	})
	v = gvar.New(result)
	err = v.Structs(&topMenus)
	return topMenus, err
}

// 根据给定的Url检索顶部菜单，给定的Url可能只是一个Url Path。
func (s *sMenu) GetTopMenuByUrl(ctx context.Context, url string) (*model.MenuItem, error) {
	items, _ := s.GetTopMenus(ctx)
	for _, v := range items {
		if v.Url == url {
			return v, nil
		}
	}
	return nil, nil
}
