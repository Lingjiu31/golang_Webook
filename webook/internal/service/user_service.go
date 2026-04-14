package service

import (
	"Project-WeBook/webook/internal/domain"
	"Project-WeBook/webook/internal/repository"
	"context"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) SignUp(cxt context.Context, user domain.User) error {
	return svc.repo.Create(cxt, user)
}
