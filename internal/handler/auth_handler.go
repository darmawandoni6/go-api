package handler

import (
	"net/http"

	"github.com/donidarmawan/test-go/internal/dto"
	"github.com/donidarmawan/test-go/internal/middleware"
	"github.com/donidarmawan/test-go/internal/repository"
	"github.com/donidarmawan/test-go/internal/service"
	"github.com/gin-gonic/gin"
)

// AuthHandler handles HTTP requests for authentication endpoints.
type AuthHandler struct {
	authService service.AuthService
	userRepo    repository.UserRepository
}

// NewAuthHandler creates a new AuthHandler.
func NewAuthHandler(authService service.AuthService, userRepo repository.UserRepository) *AuthHandler {
	return &AuthHandler{authService: authService, userRepo: userRepo}
}

// POST /api/v1/auth/register
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	user, err := h.authService.Register(&req)
	if err != nil {
		if err.Error() == "email already registered" {
			c.JSON(http.StatusConflict, dto.ErrorResponse{
				Code:    http.StatusConflict,
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "failed to register user",
		})
		return
	}

	c.JSON(http.StatusCreated, dto.SuccessResponse{
		Code:    http.StatusCreated,
		Message: "registration successful",
		Data: dto.UserResponse{
			ID:        user.ID.String(),
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		},
	})
}

// POST /api/v1/auth/login
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	resp, err := h.authService.Login(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Code:    http.StatusOK,
		Message: "login successful",
		Data:    resp,
	})
}

// GET /api/v1/auth/me  (protected by AuthGuard middleware)
func (h *AuthHandler) Me(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)

	user, err := h.userRepo.FindByID(userID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data: dto.UserResponse{
			ID:        user.ID.String(),
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		},
	})
}
