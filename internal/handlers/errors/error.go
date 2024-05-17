package errors

import (
	"basic-api/internal/response"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// HTTPErrorHandler http error handler
func HTTPErrorHandler(err error, c echo.Context) {
	var code int
	var message interface{}

	if strings.Contains(err.Error(), "pq") || strings.Contains(err.Error(), "LDAP") {
		code = http.StatusBadRequest
		message = map[string]interface{}{"message": "Internal data error"}
	} else if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = map[string]interface{}{"message": he.Message}
	} else if he, ok := err.(response.Error); ok {
		code = http.StatusBadRequest
		message = he

		if he.Code == 401 {
			code = http.StatusUnauthorized
			message = map[string]string{"message": he.Message}
		}
	} else {
		code = http.StatusBadRequest
		message = map[string]string{"message": err.Error()}
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, message)
		}

		if err != nil {
			c.Logger().Error(err)
		}
	}
}
