package menu

import (
	"basic-api/internal/app"
	"basic-api/internal/core/handler"

	"github.com/labstack/echo/v4"
)

// Interface user interface
type Interface interface {
	GetAll(c echo.Context) error
}

// Endpoint user endpoint
type Endpoint struct {
	service ServiceInterface
}

// NewEndpoint new endpoint,
func NewEndpoint(c *app.Context) Interface {
	return &Endpoint{
		service: NewService(c),
	}
}

// GetAll get all
// @Summary Get all menu
// @Description get all menu
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Accept-Language header string false "(en, th)" default(en)
// @Param request query GetAllForm true "form for get all"
// @Success 200 {object} models.Page
// @Failure 400 {object} models.Message
// @Failure 401 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 410 {object} models.Message
// @Failure 503 {object} models.Maintenance
// @Tags Menu
// @Router /menu [GET]
func (e *Endpoint) GetAll(c echo.Context) error {
	return handler.ResponseObject(c, e.service.GetAll, &GetAllForm{})
}
