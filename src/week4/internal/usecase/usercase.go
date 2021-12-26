package usecase

import (
	"context"
	"go_practice/go_practice/src/week4/internal/domain"
)

type user struct {
	repo domain.IUserRepo
}

// NewUserUsecase NewUserUsecase
func NewUserUsecase(repo domain.IUserRepo) domain.IUserUsecase {
	return &user{repo: repo}
}

// GetUserInfo GetUserInfo
func (u *user) GetUserInfo(ctx context.Context, id int) (*domain.User, error) {
	return u.repo.GetUserInfo(ctx, id)
}
