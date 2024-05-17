package menu

import (
	"basic-api/internal/entities"
)

// GetAllForm get all form
type GetAllForm struct {
	entities.PageForm
}

type StdOutModel struct {
	Status bool             `json:"status"`
	Code   int              `json:"code"`
	Data   []*entities.Menu `json:"data"`
}
