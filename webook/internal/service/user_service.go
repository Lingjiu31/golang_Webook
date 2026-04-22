package service

import (
	"Project-WeBook/webook/internal/domain"
	"Project-WeBook/webook/internal/repository"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserDuplicateEmail    = repository.ErrUserDuplicateEmail
	ErrInvalidUserOrPassword = errors.New("账号(邮箱)或密码错误")
	ErrUserNotFound          = repository.ErrUserNotFound
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) SignUp(ctx context.Context, user domain.User) error {
	//加密
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)

	return svc.repo.Create(ctx, user)
}

func (svc *UserService) Login(ctx context.Context, user domain.User) (domain.User, error) {
	// 先找用户
	dbUser, err := svc.repo.FindByEmail(ctx, user.Email)
	if errors.Is(err, ErrUserNotFound) {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	if err != nil {
		return domain.User{}, err
	}

	// 比较密码
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	return domain.User{
		Id:       dbUser.Id,
		Email:    dbUser.Email,
		Password: dbUser.Password,
	}, nil
}

func (svc *UserService) Edit(ctx context.Context, user domain.User) error {

	return svc.repo.Edit(ctx, user)
}

func (svc *UserService) GetProfile(ctx context.Context, id int64) (domain.User, error) {
	user, err := svc.repo.FindById(ctx, id)
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
