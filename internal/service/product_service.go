package service

import (
	"github.com/donidarmawan/test-go/internal/dto"
	"github.com/donidarmawan/test-go/internal/models"
	"github.com/donidarmawan/test-go/internal/repository"
)

type ProductService interface {
	CreateCategory(req *dto.CreateCategoryRequest) (*models.Category, error)
	UpdateCategory(id uint, req *dto.UpdateCategoryRequest) (*models.Category, error)
	GetCategoryById(id uint) (*models.Category, error)
	GetCategories() ([]models.Category, error)
	DeleteCategory(id uint) error
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{productRepo: productRepo}
}

func (s *productService) CreateCategory(req *dto.CreateCategoryRequest) (*models.Category, error) {
	return s.productRepo.CreateCategory(req)
}

func (s *productService) UpdateCategory(id uint, req *dto.UpdateCategoryRequest) (*models.Category, error) {
	return s.productRepo.UpdateCategory(id, req)
}

func (s *productService) GetCategoryById(id uint) (*models.Category, error) {
	return s.productRepo.GetCategoryById(id)
}

func (s *productService) GetCategories() ([]models.Category, error) {
	return s.productRepo.GetCategories()
}

func (s *productService) DeleteCategory(id uint) error {
	return s.productRepo.DeleteCategory(id)
}
