package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gmode"
	"go-blog/internal/controller"
	"go-blog/internal/controller/blog"
	"go-blog/internal/controller/i"
	"go-blog/internal/service"

	"go-blog/internal/consts"
	"go-blog/utility/response"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start focus server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			var (
				s   = g.Server()
				oai = s.GetOpenApi()
			)
			domain := g.Cfg().MustGet(ctx, "setting.domain").String()
			// OpenApi自定义信息
			oai.Info.Title = `API Reference`
			oai.Config.CommonResponse = response.JsonRes{}
			oai.Config.CommonResponseDataField = `Data`

			// 静态目录设置
			uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
			if uploadPath == "" {
				g.Log().Fatal(ctx, "文件上传配置路径不能为空")
			}
			s.AddStaticPath("/upload", uploadPath)

			//设置cookies域名方便使用二级域名
			s.SetCookieDomain("." + domain)
			// HOOK, 开发阶段禁止浏览器缓存,方便调试
			if gmode.IsDevelop() {
				s.BindHookHandler("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
					r.Response.Header().Set("Cache-Control", "no-store")
				})
			}

			// 前台系统自定义错误页面
			s.BindStatusHandler(401, func(r *ghttp.Request) {
				if !gstr.HasPrefix(r.URL.Path, "/admin") {
					service.View().Render401(r.Context())
				}
			})
			s.BindStatusHandler(403, func(r *ghttp.Request) {
				if !gstr.HasPrefix(r.URL.Path, "/admin") {
					service.View().Render403(r.Context())
				}
			})
			s.BindStatusHandler(404, func(r *ghttp.Request) {
				if !gstr.HasPrefix(r.URL.Path, "/admin") {
					service.View().Render404(r.Context())
				}
			})
			s.BindStatusHandler(500, func(r *ghttp.Request) {
				if !gstr.HasPrefix(r.URL.Path, "/admin") {
					service.View().Render500(r.Context())
				}
			})

			/**
			main
			*/
			mainDomain := s.Domain(domain)
			mainDomain.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				group.Bind(
					controller.Captcha, // 验证码
					controller.Index,   // 首页
					controller.Page,    // 首页
				)

				// 权限控制路由
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Auth)
					group.Bind(
						controller.File, // 文件
					)
				})
			})

			/**
			用户中心
			*/
			userCenter := s.Domain("i." + domain)
			userCenter.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				group.Bind(
					i.Login,    // 登录
					i.Register, // 登录
					i.User,     // 登录
				)
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Auth)
					group.Bind(
						i.Profile,       // 个人
						controller.File, // 文件
					)
				})

				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Auth)
					group.Middleware(service.Middleware().AuthAdmin)
					group.Bind(
						i.Page,
						i.Content, // 内容
					)
				})
			})

			/**
			博客
			*/
			blogCenter := s.Domain("blog." + domain)
			blogCenter.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				group.Bind(
					blog.Blog, // 文章
				)
			})
			// 自定义丰富文档
			enhanceOpenAPIDoc(s)
			// 启动Http Server
			s.Run()
			return
		},
	}
)

func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `Data`

	// API description.
	openapi.Info.Title = `Focus Project`
	openapi.Info.Description = ``

	// Sort the tags in custom sequence.
	openapi.Tags = &goai.Tags{
		{Name: consts.OpenAPITagNameIndex},
		{Name: consts.OpenAPITagNameLogin},
		{Name: consts.OpenAPITagNameRegister},
		{Name: consts.OpenAPITagNameArticle},
		{Name: consts.OpenAPITagNameTopic},
		{Name: consts.OpenAPITagNameAsk},
		{Name: consts.OpenAPITagNameReply},
		{Name: consts.OpenAPITagNameContent},
		{Name: consts.OpenAPITagNameSearch},
		{Name: consts.OpenAPITagNameInteract},
		{Name: consts.OpenAPITagNameCategory},
		{Name: consts.OpenAPITagNameProfile},
		{Name: consts.OpenAPITagNameUser},
		{Name: consts.OpenAPITagNameMess},
	}
}
