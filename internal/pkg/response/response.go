package response

import "github.com/gin-gonic/gin"

// Response 统一 API 响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

// PageResponse 分页响应
type PageResponse struct {
	List  interface{} `json:"list"`
	Total int64       `json:"total"`
	Page  int         `json:"page"`
	Size  int         `json:"size"`
}

// OK 成功响应
func OK(c *gin.Context, data interface{}) {
	c.JSON(200, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

// OKWithMeta 成功响应（带元数据）
func OKWithMeta(c *gin.Context, data, meta interface{}) {
	c.JSON(200, Response{
		Code:    0,
		Message: "success",
		Data:    data,
		Meta:    meta,
	})
}

// Page 分页成功响应
func Page(c *gin.Context, list interface{}, total int64, page, size int) {
	c.JSON(200, Response{
		Code:    0,
		Message: "success",
		Data: PageResponse{
			List:  list,
			Total: total,
			Page:  page,
			Size:  size,
		},
	})
}

// Fail 失败响应
func Fail(c *gin.Context, code int, message string) {
	c.JSON(200, Response{
		Code:    code,
		Message: message,
	})
}

// FailWithData 失败响应（带数据）
func FailWithData(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(200, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
