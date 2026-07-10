package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/dongowu/gokick/internal/pkg/apperror"
	"github.com/dongowu/gokick/internal/service/user"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

// CreateUser 创建User
// @Summary 创建User
// @Description 创建新User
// @Tags user
// @Accept json
// @Produce json
// @Param request body service.CreateUserRequest true "请求参数"
// @Success 200 {object} response.Response
// @Failure 400 {object} apperror.AppError
// @Router /user/user [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req service.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperror.Wrap(err, apperror.ErrInvalidParams.Code, "参数解析失败"))
		return
	}

	resp, err := h.svc.CreateUser(c.Request.Context(), &req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": resp})
}
