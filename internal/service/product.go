package service

import (
	"context"

	"github.com/dongowu/gokick/internal/pkg/apperror"
	"github.com/dongowu/gokick/internal/repository/product"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

type CreateProductRequest struct {
	// TODO: 添加字段
}

type CreateProductResponse struct {
	ID uint `json:"id"`
}

func (s *ProductService) CreateProduct(ctx context.Context, req *CreateProductRequest) (*CreateProductResponse, error) {
	// TODO: 实现业务逻辑
	return &CreateProductResponse{ID: 1}, nil
}
