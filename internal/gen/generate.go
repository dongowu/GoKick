package gen

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	handlerTmpl = `package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/dongowu/gokick/internal/pkg/apperror"
	"github.com/dongowu/gokick/internal/service/{{.Module}}"
)

type {{.Entity}}Handler struct {
	svc *service.{{.Entity}}Service
}

func New{{.Entity}}Handler(svc *service.{{.Entity}}Service) *{{.Entity}}Handler {
	return &{{.Entity}}Handler{svc: svc}
}

// Create{{.Entity}} 创建{{.Entity}}
// @Summary 创建{{.Entity}}
// @Description 创建新{{.Entity}}
// @Tags {{.Module}}
// @Accept json
// @Produce json
// @Param request body service.Create{{.Entity}}Request true "请求参数"
// @Success 200 {object} response.Response
// @Failure 400 {object} apperror.AppError
// @Router /{{.Module}}/{{.EntityLower}} [post]
func (h *{{.Entity}}Handler) Create{{.Entity}}(c *gin.Context) {
	var req service.Create{{.Entity}}Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperror.Wrap(err, apperror.ErrInvalidParams.Code, "参数解析失败"))
		return
	}

	resp, err := h.svc.Create{{.Entity}}(c.Request.Context(), &req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": resp})
}
`

	serviceTmpl = `package service

import (
	"context"

	"github.com/dongowu/gokick/internal/pkg/apperror"
	"github.com/dongowu/gokick/internal/repository/{{.Module}}"
)

type {{.Entity}}Service struct {
	repo *repository.{{.Entity}}Repository
}

func New{{.Entity}}Service(repo *repository.{{.Entity}}Repository) *{{.Entity}}Service {
	return &{{.Entity}}Service{repo: repo}
}

type Create{{.Entity}}Request struct {
	// TODO: 添加字段
}

type Create{{.Entity}}Response struct {
	ID uint ` + "`json:\"id\"`" + `
}

func (s *{{.Entity}}Service) Create{{.Entity}}(ctx context.Context, req *Create{{.Entity}}Request) (*Create{{.Entity}}Response, error) {
	// TODO: 实现业务逻辑
	return &Create{{.Entity}}Response{ID: 1}, nil
}
`

	repositoryTmpl = `package repository

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type {{.Entity}}Repository struct {
	db *gorm.DB
}

func New{{.Entity}}Repository(db *gorm.DB) *{{.Entity}}Repository {
	return &{{.Entity}}Repository{db: db}
}

type {{.Entity}}Model struct {
	ID        uint           ` + "`gorm:\"primaryKey\"`" + `
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt ` + "`gorm:\"index\"`" + `
	// TODO: 添加业务字段
}

func ({{.Entity}}Model) TableName() string {
	return "{{.EntityLower}}s"
}

func (r *{{.Entity}}Repository) Create(ctx context.Context, model *{{.Entity}}Model) error {
	return r.db.WithContext(ctx).Create(model).Error
}

func (r *{{.Entity}}Repository) GetByID(ctx context.Context, id uint) (*{{.Entity}}Model, error) {
	var model {{.Entity}}Model
	err := r.db.WithContext(ctx).First(&model, id).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (r *{{.Entity}}Repository) Update(ctx context.Context, model *{{.Entity}}Model) error {
	return r.db.WithContext(ctx).Save(model).Error
}

func (r *{{.Entity}}Repository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&{{.Entity}}Model{}, id).Error
}

func (r *{{.Entity}}Repository) List(ctx context.Context, offset, limit int) ([]{{.Entity}}Model, error) {
	var models []{{.Entity}}Model
	err := r.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&models).Error
	return models, err
}
`

	modelTmpl = `package model

import "time"

type {{.Entity}} struct {
	ID        uint      ` + "`json:\"id\" gorm:\"primaryKey\"`" + `
	CreatedAt time.Time ` + "`json:\"created_at\"`" + `
	UpdatedAt time.Time ` + "`json:\"updated_at\"`" + `
	// TODO: 添加业务字段
}

func ({{.Entity}}) TableName() string {
	return "{{.EntityLower}}s"
}
`
)

type TemplateData struct {
	Entity      string
	Module      string
	EntityLower string
}

func GenerateHandler(data TemplateData) (string, error) {
	return executeTemplate(handlerTmpl, data)
}

func GenerateService(data TemplateData) (string, error) {
	return executeTemplate(serviceTmpl, data)
}

func GenerateRepository(data TemplateData) (string, error) {
	return executeTemplate(repositoryTmpl, data)
}

func GenerateModel(data TemplateData) (string, error) {
	return executeTemplate(modelTmpl, data)
}

func executeTemplate(tmpl string, data TemplateData) (string, error) {
	t, err := template.New("gen").Parse(tmpl)
	if err != nil {
		return "", err
	}
	var buf strings.Builder
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func WriteFile(path, content string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return os.WriteFile(path, []byte(content), 0644)
}
