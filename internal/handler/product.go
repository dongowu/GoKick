package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/dongowu/gokick/internal/pkg/apperror"
	"github.com/dongowu/gokick/internal/service/product"
)

type ProductHandler struct {
	svc *service.ProductService
}

func NewProductHandler(svc *service.ProductService) *ProductHandler {
	return &ProductHandler{svc: svc}
}

// CreateProduct 创建Product
// @Summary 创建Product
// @Description 创建新Product
// @Tags product
// @Accept json
// @Produce json
// @Param request body service.CreateProductRequest true "请求参数"
// @Success 200 {object} response.Response
// @Failure 400 {object} apperror.AppError
// @Router /product/product [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req service.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperror.Wrap(err, apperror.ErrInvalidParams.Code, "参数解析失败"))
		return
	}

	resp, err := h.svc.CreateProduct(c.Request.Context(), &req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": resp})
}
