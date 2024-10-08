package i

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"go-blog/api/v1"
	"go-blog/internal/consts"
	"go-blog/internal/model"
	"go-blog/internal/service"
)

// 注册控制器
var Register = cRegister{}

type cRegister struct{}

func (a *cRegister) Index(ctx context.Context, req *v1.RegisterIndexReq) (res *v1.RegisterIndexRes, err error) {
	service.View().RenderTpl(ctx, "index/i/auth.html", model.View{
		MainTpl: "index/i/register/index.html",
	})
	return
}

func (a *cRegister) Register(ctx context.Context, req *v1.RegisterDoReq) (res *v1.RegisterDoRes, err error) {
	if !service.Captcha().VerifyAndClear(g.RequestFromCtx(ctx), consts.CaptchaDefaultName, req.Captcha) {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "请输入正确的验证码")
	}
	if err = service.User().Register(ctx, model.UserRegisterInput{
		Passport: req.Passport,
		Password: req.Password,
		Nickname: req.Nickname,
	}); err != nil {
		return
	}

	// 自动登录
	err = service.User().Login(ctx, model.UserLoginInput{
		Passport: req.Passport,
		Password: req.Password,
	})
	return
}
