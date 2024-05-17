package middleware

import (
	"fmt"

	"basic-api/internal/response"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	XApiKey = "bWFzdGVyLXgtYXBpLWtleQ=="
)

func ApiKeyGuard(confKey string) echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: fmt.Sprintf("header:%s", "x-api-key"),
		Validator: func(key string, c echo.Context) (bool, error) {
			return key == confKey, nil
		},
		ErrorHandler: func(err error, c echo.Context) error {
			if err != nil {
				return response.Error{Code: 401, Message: "Unauthorized"}
			}
			return nil
		},
	})
}
