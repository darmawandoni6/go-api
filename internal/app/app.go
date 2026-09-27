package app

import (
	"github.com/donidarmawan/test-go/internal/config"
	"github.com/donidarmawan/test-go/internal/handler"
	"github.com/donidarmawan/test-go/internal/repository"
	"github.com/donidarmawan/test-go/internal/router"
	"github.com/donidarmawan/test-go/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// App holds the wired-up HTTP engine ready to serve.
type App struct {
	Router *gin.Engine
}

// New wires all dependencies and registers routes, returning a runnable App.
// Add new modules here as the project grows.
func New(db *gorm.DB, cfg *config.Config) *App {
	// Auth module
	userRepo    := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo, cfg)
	authHandler := handler.NewAuthHandler(authService, userRepo)

	// Product module
	productRepo     := repository.NewProductRepository(db)
	productService  := service.NewProductService(productRepo)
	productHandler  := handler.NewProductHandler(productService)

	// Build Gin engine & register all routes
	// gin.New() instead of gin.Default() — Recovery & Logger are set explicitly
	r := gin.New()
	r.Use(gin.Logger())
	router.Setup(r, router.Handlers{
		Auth:    authHandler,
		Product: productHandler,
	}, cfg)

	return &App{Router: r}
}

