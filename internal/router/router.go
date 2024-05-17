package router

import (
	"basic-api/internal/app"
	"basic-api/internal/handlers/beer"
	"basic-api/internal/handlers/errors"
	"basic-api/internal/handlers/menu"

	// "basic-api/internal/handlers/errors"
	coreMiddleware "basic-api/internal/middleware"
	"net/http"

	"basic-api/internal/core/validator"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// Options option for new router
type Options struct {
	LogRequestFormat string
	AppContext       *app.Context
}

// New new router
func New() *echo.Echo {
	return NewWithOptions(nil)
}

// NewWithOptions new router with options
func NewWithOptions(options *Options) *echo.Echo {
	router := echo.New()
	router.Validator = validator.New()
	router.HTTPErrorHandler = errors.HTTPErrorHandler

	router.Use(
		coreMiddleware.CustomContext(options.AppContext.Db, logrus.StandardLogger()),
		coreMiddleware.ApiKeyGuard(coreMiddleware.XApiKey),
	)

	// Create Log
	// .. not implement yet

	// Endpoint endpoint
	beerEndpoint := beer.NewEndpoint(options.AppContext)
	menuEndpoint := menu.NewEndpoint(options.AppContext)

	// Base Api base api
	api := router.Group("/api")

	// Health Check health check
	healthcheck := api.Group("/healthcheck")
	{
		healthcheck.GET("", func(c echo.Context) error {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "healthy",
			})
		})
	}

	// Beer Group beer group
	beerGroup := api.Group("/beer")
	{
		beerGroup.POST("", beerEndpoint.Create)
		beerGroup.GET("", beerEndpoint.GetAll)
		beerGroup.GET("/:id", beerEndpoint.GetOne)
		beerGroup.PUT("/:id", beerEndpoint.Update)
		beerGroup.DELETE("/:id", beerEndpoint.Delete)
	}

	// Menu menu
	menuGroup := api.Group("/menu")
	{
		menuGroup.GET("", menuEndpoint.GetAll)
	}

	return router
}
