package service

import (
	"context"
	"errors"
	"go-basic-exercises/webook/internal/domain"
	"go-basic-exercises/webook/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

// @program:   go-basic-exercises
// @file:      user.go
// @author:    bowen
// @time:      2023-09-21 21:03
// @description:

var ErrEmailDuplicate = repository.ErrEmailDuplicate
var InvalidUserOrPasswd = errors.New("账号不对或者密码不对")
var ErrUserNotfound = repository.ErrUserNotfound

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {

	return &UserService{repo: repo}
}

// create
// 保持链路，链路追踪
// 注册
func (svc *UserService) SignUp(c context.Context, u domain.User) error {
	// 加密
	// 存起来
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return svc.repo.Create(c, u)
}

// 登录
func (svc *UserService) Login(c context.Context, email, password string) (domain.User, error) {
	// 找用户
	u, err := svc.repo.FindByEmail(c, email)

	if err == repository.ErrUserNotfound {
		// 这里也返回了用户或密码错误
		return domain.User{}, InvalidUserOrPasswd
	}
	if err != nil {
		return domain.User{}, err
	}
	//这里能写出来吗，比较密码
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	//返回密码输入错误
	if err != nil {
		//大日志
		return domain.User{}, InvalidUserOrPasswd
	}

	return u, nil
}
