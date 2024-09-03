package i

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"go-blog/api/v1"
	"go-blog/internal/consts"
	"go-blog/internal/service"
)

// 用户管理
var User = cUser{}

type cUser struct{}

func (a *cUser) Logout(ctx context.Context, req *v1.UserLogoutReq) (res *v1.UserLogoutRes, err error) {
	if err = service.User().Logout(ctx); err != nil {
		return
	}
	g.RequestFromCtx(ctx).Response.RedirectTo(consts.UserLoginUrl)
	return
}
