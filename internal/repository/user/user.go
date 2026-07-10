package repository

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

type UserModel struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	// TODO: 添加业务字段
}

func (UserModel) TableName() string {
	return "users"
}

func (r *UserRepository) Create(ctx context.Context, model *UserModel) error {
	return r.db.WithContext(ctx).Create(model).Error
}

func (r *UserRepository) GetByID(ctx context.Context, id uint) (*UserModel, error) {
	var model UserModel
	err := r.db.WithContext(ctx).First(&model, id).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (r *UserRepository) Update(ctx context.Context, model *UserModel) error {
	return r.db.WithContext(ctx).Save(model).Error
}

func (r *UserRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&UserModel{}, id).Error
}

func (r *UserRepository) List(ctx context.Context, offset, limit int) ([]UserModel, error) {
	var models []UserModel
	err := r.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&models).Error
	return models, err
}
