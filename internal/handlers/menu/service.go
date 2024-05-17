package menu

import (
	"basic-api/internal/app"
	"basic-api/internal/core/context"
)

// ServiceInterface service interface
type ServiceInterface interface {
	GetAll(c *context.Context, f *GetAllForm) (interface{}, error)
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

// GetAll get all
func (s *Service) GetAll(c *context.Context, f *GetAllForm) (interface{}, error) {
	objs, err := s.repository.FindAllMenu(c.Db, f)
	if err != nil {
		c.Log.Error(err)
		return nil, err
	}

	return StdOutModel{
		Status: true,
		Code:   200,
		Data:   objs,
	}, nil
}
