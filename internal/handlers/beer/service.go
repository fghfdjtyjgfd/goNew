package beer

import (
	"basic-api/internal/app"
	"basic-api/internal/core/context"
	"basic-api/internal/core/form"
	"basic-api/internal/entities"
	"log"

	"github.com/jinzhu/copier"
)

// ServiceInterface service interface
type ServiceInterface interface {
	Create(c *context.Context, f *CreateForm) (*entities.Beer, error)
	GetOne(c *context.Context, f *form.GetOneForm) (*entities.Beer, error)
	GetAll(c *context.Context, f *GetAllForm) (*entities.Page, error)
	Update(c *context.Context, f *UpdateForm) error
	Delete(c *context.Context, f *form.GetOneForm) error
}

// Service service
type Service struct {
	repository RepoInterface
}

// NewService new service
func NewService(c *app.Context) *Service {
	return &Service{
		repository: NewRepo(c),
	}
}

// Create create
func (s *Service) Create(c *context.Context, f *CreateForm) (*entities.Beer, error) {
	object := &entities.Beer{}
	_ = copier.Copy(object, f)

	err := s.repository.Create(c.Db, object)
	if err != nil {
		c.Log.Error(err)
		return nil, err
	}

	return object, nil
}

// GetOne get one
func (s *Service) GetOne(c *context.Context, f *form.GetOneForm) (*entities.Beer, error) {
	object := &entities.Beer{}
	err := s.repository.FindOneObjectByID(c.Db, f.ID, object)
	if err != nil {
		c.Log.Error(err)
		return nil, err
	}

	return object, nil
}

// GetAll get all
func (s *Service) GetAll(c *context.Context, f *GetAllForm) (*entities.Page, error) {
	log.Println(f.Query)
	page, err := s.repository.FindAllByQuery(c.Db, f)
	if err != nil {
		c.Log.Error(err)
		return nil, err
	}

	return page, nil
}

// Update update
func (s *Service) Update(c *context.Context, f *UpdateForm) error {
	object := &entities.Beer{}
	err := s.repository.FindOneObjectByID(c.Db, f.ID, object)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	_ = copier.Copy(object, f)
	err = s.repository.Update(c.Db, object)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	return nil
}

// Delete delete
func (s *Service) Delete(c *context.Context, f *form.GetOneForm) error {
	object := &entities.Beer{}
	err := s.repository.FindOneObjectByID(c.Db, f.ID, object)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	err = s.repository.Delete(c.Db, object)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	return nil
}
