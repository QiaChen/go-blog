package i

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"go-blog/api/v1"
	"go-blog/internal/model"
	"go-blog/internal/service"
	"runtime"
)

// 个人中心
var Profile = cProfile{}

type cProfile struct{}

type systemInfo struct {
	NumCPU       int
	GOMAXPROCS   int
	Version      string
	Mem          string
	NumGoroutine int
}

func (a *cProfile) Dash(ctx context.Context, req *v1.ProfileDashReq) (res *v1.ProfileDashRes, err error) {
	title := "用户中心"
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	out := systemInfo{
		Version:      runtime.Version(),
		NumCPU:       runtime.NumCPU(),
		GOMAXPROCS:   runtime.GOMAXPROCS(0),
		NumGoroutine: runtime.NumGoroutine(),
		Mem:          fmt.Sprintf("Alloc: %vMB, TotalAlloc: %vMB, Sys: %vMB, NumGC: %v\n", m.Alloc/1024/1024, m.TotalAlloc/1024/1024, m.Sys/1024/1024, m.NumGC),
	}

	service.View().RenderTpl(ctx, "index/i/user_layout.html", model.View{
		MainTpl:     "index/i/index.html",
		Title:       title,
		Keywords:    title,
		Description: title,
		Data:        out,
	})
	return nil, nil
}

func (a *cProfile) Index(ctx context.Context, req *v1.ProfileIndexReq) (res *v1.ProfileIndexRes, err error) {
	out, err := service.User().GetProfile(ctx)
	if err != nil {
		return nil, err
	}
	title := "用户 " + out.Nickname + " 资料"
	service.View().RenderTpl(ctx, "index/i/user_layout.html", model.View{
		MainTpl:     "index/i/profile/index.html",
		Title:       title,
		Keywords:    title,
		Description: title,
		Data:        out,
	})
	return nil, nil
}

func (a *cProfile) Avatar(ctx context.Context, req *v1.ProfileAvatarReq) (res *v1.ProfileAvatarRes, err error) {
	out, err := service.User().GetProfile(ctx)
	if err != nil {
		return nil, err
	}
	title := "用户 " + out.Nickname + " 头像"
	service.View().RenderTpl(ctx, "index/i/user_layout.html", model.View{
		MainTpl:     "index/i/profile/avatar.html",
		Title:       title,
		Keywords:    title,
		Description: title,
		Data:        out,
	})
	return nil, nil
}

func (a *cProfile) UpdateAvatar(ctx context.Context, req *v1.ProfileUpdateAvatarReq) (res *v1.ProfileUpdateAvatarRes, err error) {
	var (
		request    = g.RequestFromCtx(ctx)
		uploadFile = request.GetUploadFile("file")
	)
	if uploadFile == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请选择需要上传的文件")
	}
	uploadResult, err := service.File().Upload(ctx, model.FileUploadInput{
		File:       uploadFile,
		RandomName: true,
	})
	if err != nil {
		return nil, err
	}
	if uploadResult != nil {
		req.Avatar = uploadResult.Url
	}
	err = service.User().UpdateAvatar(ctx, model.UserUpdateAvatarInput{
		UserId: service.BizCtx().Get(ctx).User.Id,
		Avatar: req.Avatar,
	})
	if err != nil {
		return nil, err
	}
	// 更新登录session Avatar
	sessionProfile := service.Session().GetUser(ctx)
	sessionProfile.Avatar = req.Avatar
	err = service.Session().SetUser(ctx, sessionProfile)
	return
}

func (a *cProfile) Password(ctx context.Context, req *v1.ProfilePasswordReq) (res *v1.ProfilePasswordRes, err error) {
	out, err := service.User().GetProfile(ctx)
	if err != nil {
		return nil, err
	}
	title := "用户 " + out.Nickname + " 修改密码"
	service.View().RenderTpl(ctx, "index/i/user_layout.html", model.View{
		MainTpl:     "index/i/profile/password.html",
		Title:       title,
		Keywords:    title,
		Description: title,
		Data:        out,
	})
	return nil, nil
}

func (a *cProfile) UpdatePassword(ctx context.Context, req *v1.ProfileUpdatePasswordReq) (res *v1.ProfileUpdatePasswordRes, err error) {
	err = service.User().UpdatePassword(ctx, model.UserPasswordInput{
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})
	return
}

func (a *cProfile) UpdateProfile(ctx context.Context, req *v1.ProfileUpdateReq) (res *v1.ProfileUpdateRes, err error) {
	err = service.User().UpdateProfile(ctx, model.UserUpdateProfileInput{
		UserId:   req.Id,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Gender:   req.Gender,
	})
	return
}
