package beer

import (
	"basic-api/internal/app"
	"basic-api/internal/core/form"
	"basic-api/internal/core/handler"

	"github.com/labstack/echo/v4"
)

// Interface user interface
type Interface interface {
	Create(c echo.Context) error
	GetAll(c echo.Context) error
	GetOne(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
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

// Create create
// @Summary Create beer
// @Description create beer
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Accept-Language header string false "(en, th)" default(en)
// @Param request body CreateForm true "form for create"
// @Success 200 {object} models.AlbumTopic
// @Failure 400 {object} models.Message
// @Failure 401 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 410 {object} models.Message
// @Failure 503 {object} models.Maintenance
// @Tags Beer
// @Router /beer [POST]
func (e *Endpoint) Create(c echo.Context) error {
	return handler.ResponseObject(c, e.service.Create, &CreateForm{})
}

// GetAll get all
// @Summary Get all beer
// @Description get all beer
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
// @Tags Beer
// @Router /beer [GET]
func (e *Endpoint) GetAll(c echo.Context) error {
	return handler.ResponseObject(c, e.service.GetAll, &GetAllForm{})
}

// GetOne get one
// @Summary Get beer
// @Description get beer
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Accept-Language header string false "(en, th)" default(en)
// @Param id path int true "ID"
// @Success 200 {object} models.AlbumTopic
// @Failure 400 {object} models.Message
// @Failure 401 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 410 {object} models.Message
// @Failure 503 {object} models.Maintenance
// @Tags Beer
// @Router /beer/{id} [GET]
func (e *Endpoint) GetOne(c echo.Context) error {
	return handler.ResponseObject(c, e.service.GetOne, &form.GetOneForm{})
}

// Update update
// @Summary Update beer
// @Description update beer
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Accept-Language header string false "(en, th)" default(en)
// @Param id path int true "ID"
// @Param request body UpdateForm true "form for update"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Message
// @Failure 401 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 410 {object} models.Message
// @Failure 503 {object} models.Maintenance
// @Tags Beer
// @Router /beer/{id} [PUT]
func (e *Endpoint) Update(c echo.Context) error {
	return handler.ResponseSuccess(c, e.service.Update, &UpdateForm{})
}

// Delete delete
// @Summary Delete beer
// @Description delete beer
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Accept-Language header string false "(en, th)" default(en)
// @Param id path int true "ID"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Message
// @Failure 401 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 410 {object} models.Message
// @Failure 503 {object} models.Maintenance
// @Tags Beer
// @Router /beer/{id} [DELETE]
func (e *Endpoint) Delete(c echo.Context) error {
	return handler.ResponseSuccess(c, e.service.Delete, &form.GetOneForm{})
}
