package repo

import (
	"math"

	"gorm.io/gorm"
)

// PageFormInterface page info
type PageFormInterface interface {
	GetPage() int
	GetSize() int
	GetQuery() string
	GetSort() string
}

// Repo repo struct
type Repo struct {
}

// FindAllAndPageInformation get page information
func (r *Repo) FindAllAndPageInformation(db *gorm.DB, pf PageFormInterface, dataSource interface{}) (*PageInformation, error) {
	page := pf.GetPage()
	if pf.GetPage() < 1 {
		page = 1
	}

	limit := pf.GetSize()
	if pf.GetSize() == 0 {
		limit = 20
	}

	var count int64
	offset := 0

	db.Model(dataSource).Count(&count)
	db = db.Order("updated_at")

	if page != 1 {
		offset = (page - 1) * limit
	}
	response := db.Limit(limit).Offset(offset).Find(dataSource)
	println(response.RowsAffected)

	if response.Error != nil {
		return nil, response.Error
	}

	return &PageInformation{
		Page:                  page,
		Size:                  limit,
		TotalNumberOfEntities: count,
		TotalNumberOfPages:    int(math.Ceil(float64(count) / float64(limit))),
	}, nil
}

// FindOneObjectByID find one
func (r *Repo) FindOneObjectByID(db *gorm.DB, id uint, i interface{}) error {
	return r.FindOneObjectByField(db, "id", id, i)
}

// FindOneObjectByField find one
func (r *Repo) FindOneObjectByField(db *gorm.DB, field string, value interface{}, i interface{}) error {
	if err := db.Where(field+" = ?", value).First(i).Error; err != nil {
		return err
	}
	return nil
}

// Create create
func (r *Repo) Create(db *gorm.DB, i interface{}) error {
	if err := db.
		Set("gorm:association_autoupdate", false).
		Set("gorm:association_autocreate", false).
		Create(i).Error; err != nil {
		return err
	}
	return nil
}

// Update update
func (r *Repo) Update(db *gorm.DB, i interface{}) error {
	if err := db.
		Set("gorm:association_autoupdate", false).
		Set("gorm:association_autocreate", false).
		Save(i).Error; err != nil {
		return err
	}
	return nil
}

// Delete update
func (r *Repo) Delete(db *gorm.DB, i interface{}) error {
	if err := db.
		Set("gorm:association_autoupdate", false).
		Set("gorm:association_autocreate", false).
		Delete(i).Error; err != nil {
		return err
	}
	return nil
}
