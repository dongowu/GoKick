package service

import (
	"context"

	"github.com/dongowu/gokick/internal/model"
	"github.com/dongowu/gokick/internal/repository/user"
)

type UserService struct {
	repo *user.Repository
}

func NewUserService(repo *user.Repository) *UserService {
	return &UserService{repo: repo}
}

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email,max=100"`
	Password string `json:"password" validate:"required,min=6,max=32"`
	Nickname string `json:"nickname" validate:"max=50"`
}

type UpdateUserRequest struct {
	Nickname *string `json:"nickname" validate:"max=50"`
	Avatar   *string `json:"avatar"`
	Status   *int8   `json:"status" validate:"oneof=0 1 2"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Nickname  string    `json:"nickname"`
	Avatar    string    `json:"avatar"`
	Status    int8      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func toResponse(u *model.User) *UserResponse {
	if u == nil {
		return nil
	}
	return &UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Nickname:  u.Nickname,
		Avatar:    u.Avatar,
		Status:    u.Status,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *CreateUserRequest) (*UserResponse, error) {
	// TODO: 密码哈希
	// TODO: 校验用户名/邮箱是否已存在
	userModel := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Nickname: req.Nickname,
		Status:   1, // 默认启用
	}

	if err := s.repo.Create(ctx, userModel); err != nil {
		return nil, err
	}

	return toResponse(userModel), nil
}

func (s *UserService) GetUserByID(ctx context.Context, id uint) (*UserResponse, error) {
	u, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return toResponse(u), nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*UserResponse, error) {
	u, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return toResponse(u), nil
}

func (s *UserService) UpdateUser(ctx context.Context, id uint, req *UpdateUserRequest) (*UserResponse, error) {
	u, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// 更新字段
	if req.Nickname != nil {
		u.Nickname = *req.Nickname
	}
	if req.Avatar != nil {
		u.Avatar = *req.Avatar
	}
	if req.Status != nil {
		u.Status = *req.Status
	}

	if err := s.repo.Update(ctx, u); err != nil {
		return nil, err
	}

	return toResponse(u), nil
}

func (s *UserService) DeleteUser(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *UserService) ListUsers(ctx context.Context, offset, limit int) ([]*UserResponse, error) {
	users, err := s.repo.List(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	resp := make([]*UserResponse, 0, len(users))
	for _, u := range users {
		resp = append(resp, toResponse(u))
	}
	return resp, nil
}
