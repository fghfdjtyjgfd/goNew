package menu

import (
	"basic-api/internal/app"
	"basic-api/internal/core/repo"
	"basic-api/internal/entities"

	"gorm.io/gorm"
)

// Social social
type Social string

// RepoInterface repo interface
type RepoInterface interface {
	FindAllMenu(db *gorm.DB, f *GetAllForm) ([]*entities.Menu, error)
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

// FindAllMenu find all menu
func (r *Repo) FindAllMenu(db *gorm.DB, f *GetAllForm) ([]*entities.Menu, error) {
	db = db.Preload("Menu", func(db *gorm.DB) *gorm.DB {
		return db.Where("is_children = true")
	})

	db = db.Where("is_children = false")
	var objects []*entities.Menu

	err := db.Find(&objects).Error
	if err != nil {
		return nil, err
	}

	return objects, nil
}
