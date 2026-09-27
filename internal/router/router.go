package router

import (
	"github.com/donidarmawan/test-go/internal/config"
	"github.com/donidarmawan/test-go/internal/handler"
	"github.com/donidarmawan/test-go/internal/middleware"
	"github.com/gin-gonic/gin"
)

// Handlers bundles all HTTP handlers passed into the router.
// Add new handlers here as the project grows.
type Handlers struct {
	Auth    *handler.AuthHandler
	Product *handler.ProductHandler
}

// Setup registers all application routes onto the given Gin engine.
func Setup(r *gin.Engine, h Handlers, cfg *config.Config) {
	// Global middleware
	r.Use(middleware.Recovery())

	// Global error handlers
	r.NoRoute(middleware.NoRoute())
	r.NoMethod(middleware.NoMethod())

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// API v1 routes
	v1 := r.Group("/api-v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", h.Auth.Register)
			auth.POST("/login", h.Auth.Login)

			// Protected routes — require valid JWT
			protected := auth.Group("/")
			protected.Use(middleware.AuthGuard(cfg))
			{
				protected.GET("/me", h.Auth.Me)
			}
		}
	}

	// Category routes
	categories := v1.Group("/categories")
	categories.Use(middleware.AuthGuard(cfg))
	{
		categories.GET("", h.Product.GetCategories)
		categories.GET("/:id", h.Product.GetCategoryById)
		categories.POST("", h.Product.CreateCategory)
		categories.PUT("/:id", h.Product.UpdateCategory)
		categories.DELETE("/:id", h.Product.DeleteCategory)
	}
}
