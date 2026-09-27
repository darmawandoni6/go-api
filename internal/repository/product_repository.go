package repository

import (
	"github.com/donidarmawan/test-go/internal/dto"
	"github.com/donidarmawan/test-go/internal/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateCategory(req *dto.CreateCategoryRequest) (*models.Category, error)
	UpdateCategory(id uint, req *dto.UpdateCategoryRequest) (*models.Category, error)
	GetCategoryById(id uint) (*models.Category, error)
	GetCategories() ([]models.Category, error)
	DeleteCategory(id uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) CreateCategory(req *dto.CreateCategoryRequest) (*models.Category, error) {
	var category models.Category
	if err := r.db.Create(req).Scan(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

// DeleteCategory implements [ProductRepository].
func (r *productRepository) DeleteCategory(id uint) error {
	return r.db.Delete(&models.Category{}, id).Error
}

// GetCategories implements [ProductRepository].
func (r *productRepository) GetCategories() ([]models.Category, error) {
	var categories []models.Category
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

// GetCategoryById implements [ProductRepository].
func (r *productRepository) GetCategoryById(id uint) (*models.Category, error) {
	var category models.Category
	if err := r.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

// UpdateCategory implements [ProductRepository].
func (r *productRepository) UpdateCategory(id uint, req *dto.UpdateCategoryRequest) (*models.Category, error) {
	var category models.Category
	if err := r.db.Model(&models.Category{}).Where("id = ?", id).Updates(req).Scan(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}
