package repository

import (
	"Project-WeBook/webook/internal/domain"
	"Project-WeBook/webook/internal/repository/dao"
	"context"
)

var (
	ErrUserDuplicateEmail = dao.ErrUserDuplicateEmail
	ErrUserNotFound       = dao.ErrUserNotFound
)

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{dao: dao}
}

// Create 创建用户, 传入 domain 中结构体
// 记录到 dao 中结构体
// 创建和修改时间在 dao 中解决
func (r *UserRepository) Create(ctx context.Context, user domain.User) error {
	return r.dao.Insert(ctx, dao.User{
		Id:       user.Id,
		Email:    user.Email,
		Password: user.Password,
	})
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	user, err := r.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Id:       user.Id,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

func (r *UserRepository) Edit(ctx context.Context, user domain.User) error {
	return r.dao.UpDateUser(ctx, dao.User{
		Id:        user.Id,
		Name:      user.Name,
		Birthday:  user.Birthday,
		Biography: user.Biography,
	})
}

func (r *UserRepository) FindById(ctx context.Context, id int64) (domain.User, error) {
	user, err := r.dao.FindByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Name:      user.Name,
		Email:     user.Email,
		Birthday:  user.Birthday,
		Biography: user.Biography,
	}, nil
}
