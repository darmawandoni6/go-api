package handler

import (
	"net/http"
	"strconv"

	"github.com/donidarmawan/test-go/internal/dto"
	"github.com/donidarmawan/test-go/internal/service"
	"github.com/gin-gonic/gin"
)

// ProductHandler handles HTTP requests for product-related endpoints.
type ProductHandler struct {
	productService service.ProductService
}

// NewProductHandler creates a new ProductHandler.
func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

// parseID parses a uint path param, returning false and writing a 400 if invalid.
func parseID(c *gin.Context, param string) (uint, bool) {
	val, err := strconv.ParseUint(c.Param(param), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid " + param,
		})
		return 0, false
	}
	return uint(val), true
}

// POST /api/v1/categories
func (h *ProductHandler) CreateCategory(c *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	category, err := h.productService.CreateCategory(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "failed to create category",
		})
		return
	}

	c.JSON(http.StatusCreated, dto.SuccessResponse{
		Code:    http.StatusCreated,
		Message: "category created successfully",
		Data:    category,
	})
}

// PUT /api/v1/categories/:id
func (h *ProductHandler) UpdateCategory(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}

	var req dto.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	category, err := h.productService.UpdateCategory(id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "failed to update category",
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Code:    http.StatusOK,
		Message: "category updated successfully",
		Data:    category,
	})
}

// GET /api/v1/categories/:id
func (h *ProductHandler) GetCategoryById(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}

	category, err := h.productService.GetCategoryById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "category not found",
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    category,
	})
}

// GET /api/v1/categories
func (h *ProductHandler) GetCategories(c *gin.Context) {
	categories, err := h.productService.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "failed to get categories",
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    categories,
	})
}

// DELETE /api/v1/categories/:id
func (h *ProductHandler) DeleteCategory(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}

	if err := h.productService.DeleteCategory(id); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "failed to delete category",
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Code:    http.StatusOK,
		Message: "category deleted successfully",
	})
}
