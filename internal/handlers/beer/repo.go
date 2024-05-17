package beer

import (
	"basic-api/internal/app"
	"basic-api/internal/core/repo"
	"basic-api/internal/entities"
	"fmt"

	"gorm.io/gorm"
)

// Social social
type Social string

// RepoInterface repo interface
type RepoInterface interface {
	Create(db *gorm.DB, i interface{}) error
	Update(db *gorm.DB, i interface{}) error
	Delete(db *gorm.DB, i interface{}) error
	FindOneObjectByID(db *gorm.DB, id uint, i interface{}) error
	FindAllByQuery(db *gorm.DB, f *GetAllForm) (*entities.Page, error)
}

// Repo user repo
type Repo struct {
	repo.Repo
}

// NewRepo new service
func NewRepo(c *app.Context) *Repo {
	return &Repo{
		repo.Repo{},
	}
}

// FindAllByQuery find all by query
func (r *Repo) FindAllByQuery(db *gorm.DB, f *GetAllForm) (*entities.Page, error) {
	if f.Query != "" {
		query := f.GetQueryLike()
		db = db.Where(`name LIKE ? or description LIKE ?`, query, query)
	}

	if f.Name != nil {
		db = db.Where(`name LIKE ?`, fmt.Sprintf("%%%s%%", *f.Name))
	}

	var objects []*entities.Beer
	page, err := r.FindAllAndPageInformation(db, &f.PageForm, &objects)
	if err != nil {
		return nil, err
	}

	return entities.NewPage(page, objects), nil
}
