package beer

import (
	"basic-api/internal/core/form"
	"basic-api/internal/entities"
)

// Create From create form
type CreateForm struct {
	Name        string `json:"name" validate:"required"`
	Type        uint   `json:"type" validate:"required,min=1,max=6" swaggertype:"integer" exemple:"1"`
	Description string `json:"description" validate:"required"`
	Image       string `json:"image" validate:"required"`
}

// Update Form update form
type UpdateForm struct {
	form.GetOneForm
	CreateForm
}

// GetAllForm get all form
type GetAllForm struct {
	entities.PageForm
	Name *string `json:"name" query:"name"`
	Type *int    `json:"type" query:"type"`
}
