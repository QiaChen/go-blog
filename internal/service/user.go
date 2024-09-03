// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"go-blog/internal/model"
	"go-blog/internal/model/entity"
)

type (
	IUser interface {
		// 获得头像上传路径
		GetAvatarUploadPath() string
		// 获得头像上传对应的URL前缀
		GetAvatarUploadUrlPrefix() string
		// 执行登录
		Login(ctx context.Context, in model.UserLoginInput) error
		// 注销
		Logout(ctx context.Context) error
		// 将密码按照内部算法进行加密
		EncryptPassword(passport string, password string) string
		// 根据账号和密码查询用户信息，一般用于账号密码登录。
		// 注意password参数传入的是按照相同加密算法加密过后的密码字符串。
		GetUserByPassportAndPassword(ctx context.Context, passport string, password string) (user *entity.User, err error)
		// 检测给定的账号是否唯一
		CheckPassportUnique(ctx context.Context, passport string) error
		// 检测给定的昵称是否唯一
		CheckNicknameUnique(ctx context.Context, nickname string) error
		// 用户注册。
		Register(ctx context.Context, in model.UserRegisterInput) error
		// 修改个人密码
		UpdatePassword(ctx context.Context, in model.UserPasswordInput) error
		// 获取个人信息
		GetProfileById(ctx context.Context, userId uint) (out *model.UserGetProfileOutput, err error)
		// 修改个人资料
		GetProfile(ctx context.Context) (*model.UserGetProfileOutput, error)
		UpdateAvatar(ctx context.Context, in model.UserUpdateAvatarInput) error
		// 修改个人资料
		UpdateProfile(ctx context.Context, in model.UserUpdateProfileInput) error
		// 禁用指定用户
		Disable(ctx context.Context, id uint) error
		// 查询用户内容列表及用户信息
		GetList(ctx context.Context, in model.UserGetContentListInput) (out *model.UserGetListOutput, err error)
		// 获取文章数量
		GetUserStats(ctx context.Context, userId uint) (map[string]int, error)
		// 当前用户是否管理员
		IsCtxAdmin(ctx context.Context) bool
		// 判断给定用户是否管理员
		IsAdmin(ctx context.Context, userId uint) bool
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
