package repository

import (
	"context"
	"go-basic-exercises/webook/internal/domain"
	"go-basic-exercises/webook/internal/repository/dao"
)

// @program:   go-basic-exercises
// @file:      user.go
// @author:    bowen
// @time:      2023-09-21 21:23
// @description:

var ErrEmailDuplicate = dao.ErrEmailDuplicate
var ErrUserNotfound = dao.ErrUserNotfound

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepostitory(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao,
	}
}

func (r *UserRepository) Create(c context.Context, u domain.User) error {
	return r.dao.Insert(c, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
}

func (r *UserRepository) FindByEmail(c context.Context, email string) (domain.User, error) {

	u, err := r.dao.FindByEmail(c, email)
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:       u.Id,
		Email:    u.Email,
		Password: u.Password,
	}, nil
}
