package repository

import (
	"Project-WeBook/webook/internal/domain"
	"Project-WeBook/webook/internal/repository/dao"
	"context"
)

var ErrUserDuplicateEmail = dao.ErrUserDuplicateEmail

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{dao: dao}
}

// Create 创建用户, 传入 domain 中结构体
// 记录到 dao 中结构体
// 创建和修改时间在 dao 中解决
func (r *UserRepository) Create(cxt context.Context, user domain.User) error {
	return r.dao.Insert(cxt, dao.User{
		Email:    user.Email,
		Password: user.Password,
	})
}
