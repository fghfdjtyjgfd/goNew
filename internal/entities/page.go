package entities

import (
	"fmt"

	"basic-api/internal/core/repo"
)

// PageInformation page information
type PageInformation struct {
	Page                  int   `json:"page,omitempty"`
	Size                  int   `json:"size,omitempty"`
	TotalNumberOfEntities int64 `json:"total_number_of_entities,omitempty"`
	TotalNumberOfPages    int   `json:"total_number_of_pages,omitempty"`
}

// Page page model
type Page struct {
	PageInformation *PageInformation `json:"page_information,omitempty"`
	Entities        interface{}      `json:"entities,omitempty"`
}

// NewPage new page
func NewPage(pif *repo.PageInformation, es interface{}) *Page {
	return &Page{
		PageInformation: &PageInformation{
			Page:                  pif.Page,
			Size:                  pif.Size,
			TotalNumberOfEntities: pif.TotalNumberOfEntities,
			TotalNumberOfPages:    pif.TotalNumberOfPages,
		},
		Entities: es,
	}
}

// GetEntities get entities
func (p *Page) GetEntities() interface{} {
	return p.Entities
}

// PageForm page form
type PageForm struct {
	Page  int    `json:"page,omitempty" form:"page" query:"page"`
	Size  int    `json:"size,omitempty" form:"size" query:"size"`
	Query string `json:"query,omitempty" form:"query" query:"query"`
	Sort  string `json:"sort,omitempty" form:"sort" validate:"sort"`
}

// GetPage get page
func (f *PageForm) GetPage() int {
	return f.Page
}

// GetSize get size
func (f *PageForm) GetSize() int {
	return f.Size
}

// GetQuery get query
func (f *PageForm) GetQuery() string {
	return f.Query
}

// GetQueryLike get query like
func (f *PageForm) GetQueryLike() string {
	return fmt.Sprintf("%%%s%%", f.Query)
}

// GetSort get sort
func (f *PageForm) GetSort() string {
	return f.Sort
}
