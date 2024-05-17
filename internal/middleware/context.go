package middleware

import (
	"basic-api/internal/core/context"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// CustomContext custom context
func CustomContext(db *gorm.DB, log *logrus.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			cc := &context.Context{
				Context: c,
				Db:      db,
				Log:     log.WithField("id", ""),
			}
			return next(cc)
		}
	}
}
