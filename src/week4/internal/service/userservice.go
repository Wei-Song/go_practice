package service

import (
	"context"
	"go_practice/go_practice/src/week4/internal/domain"
	"strconv"

	v1 "go_practice/go_practice/src/week4/apis/shop/v1"

	"go_practice/go_practice/src/week4/errors/errorcode"

	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
)

// UserService UserService
type UserService struct {
	v1.UserServerServer

	usecase domain.IUserUsecase
}

// NewUserService NewUserService
func NewUserService(usecase domain.IUserUsecase) *UserService {
	return &UserService{usecase: usecase}
}

// GetUserInfo 获取用户信息
func (u *UserService) GetUserInfo(ctx context.Context, req *v1.GetUserInfoRequest) (*v1.GetUserInfoResponse, error) {
	// TODO: 下面这些应该放到中间件中
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.Wrap(errorcode.ErrUnKnown, "get metadata err")
	}

	data := md.Get("uid")
	if len(data) != 1 {
		return nil, errors.Wrapf(errorcode.ErrUnKnown, "user id lens not 1, metadata: %v", data)
	}

	id, err := strconv.Atoi(data[0])
	if err != nil {
		return nil, errors.Wrapf(errorcode.ErrUnKnown, "user id not a num, data: %v", data)
	}

	user, err := u.usecase.GetUserInfo(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := &v1.GetUserInfoResponse{
		Username: user.Name,
		City:     user.City,
	}

	return resp, nil
}
