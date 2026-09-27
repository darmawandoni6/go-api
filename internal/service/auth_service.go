package service

import (
	"errors"

	"github.com/donidarmawan/test-go/internal/config"
	"github.com/donidarmawan/test-go/internal/dto"
	"github.com/donidarmawan/test-go/internal/models"
	"github.com/donidarmawan/test-go/internal/repository"
	"github.com/donidarmawan/test-go/internal/utils"
	"gorm.io/gorm"
)

// AuthService defines the business-logic interface for authentication.
type AuthService interface {
	Register(req *dto.RegisterRequest) (*models.User, error)
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
}

type authService struct {
	userRepo repository.UserRepository
	cfg      *config.Config
}

// NewAuthService creates a new AuthService.
func NewAuthService(userRepo repository.UserRepository, cfg *config.Config) AuthService {
	return &authService{userRepo: userRepo, cfg: cfg}
}

// Register creates a new user after validating uniqueness and hashing the password.
func (s *authService) Register(req *dto.RegisterRequest) (*models.User, error) {
	// Check if email is already taken
	_, err := s.userRepo.FindByEmail(req.Email)
	if err == nil {
		return nil, errors.New("email already registered")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Hash password
	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashed,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// Login validates credentials and returns a JWT access token.
func (s *authService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !utils.CheckPassword(req.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(user.ID, user.Email, s.cfg.JWTSecret, s.cfg.JWTExpiryMinutes)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   s.cfg.JWTExpiryMinutes * 60,
	}, nil
}
