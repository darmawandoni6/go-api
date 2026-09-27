package dto

// CreateCategoryRequest holds the data needed to create a new category.
type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required,min=2"`
}

// UpdateCategoryRequest holds the data needed to update an existing category.
type UpdateCategoryRequest struct {
	Name string `json:"name" binding:"required,min=2"`
}

// CategoryResponse is the category payload returned to the client.
type CategoryResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
