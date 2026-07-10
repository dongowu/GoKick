package repository

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

type ProductModel struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	// TODO: 添加业务字段
}

func (ProductModel) TableName() string {
	return "products"
}

func (r *ProductRepository) Create(ctx context.Context, model *ProductModel) error {
	return r.db.WithContext(ctx).Create(model).Error
}

func (r *ProductRepository) GetByID(ctx context.Context, id uint) (*ProductModel, error) {
	var model ProductModel
	err := r.db.WithContext(ctx).First(&model, id).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (r *ProductRepository) Update(ctx context.Context, model *ProductModel) error {
	return r.db.WithContext(ctx).Save(model).Error
}

func (r *ProductRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&ProductModel{}, id).Error
}

func (r *ProductRepository) List(ctx context.Context, offset, limit int) ([]ProductModel, error) {
	var models []ProductModel
	err := r.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&models).Error
	return models, err
}
