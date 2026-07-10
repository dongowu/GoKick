package service

import (
	"context"

	"github.com/dongowu/gokick/internal/pkg/apperror"
	"github.com/dongowu/gokick/internal/repository/user"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

type CreateUserRequest struct {
	// TODO: 添加字段
}

type CreateUserResponse struct {
	ID uint `json:"id"`
}

func (s *UserService) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	// TODO: 实现业务逻辑
	return &CreateUserResponse{ID: 1}, nil
}
