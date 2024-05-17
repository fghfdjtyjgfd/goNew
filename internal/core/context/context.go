package context

import (
	"reflect"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	pathKey     = "path"
	userContext = "user"
)

// Parameters parameters
type Parameters interface{}

// Context custom echo context
type Context struct {
	echo.Context
	Lang       string
	Db         *gorm.DB
	Parameters Parameters
	Log        *logrus.Entry
}

// BindAndValidate bind and validate form
func (c *Context) BindAndValidate(i interface{}) error {
	if err := c.Bind(i); err != nil {
		return err
	}
	c.parsePathParams(i)
	c.Parameters = i
	if err := c.Validate(i); err != nil {
		return err
	}
	return nil
}

func (c *Context) parsePathParams(form interface{}) {
	formValue := reflect.ValueOf(form)
	if formValue.Kind() == reflect.Ptr {
		formValue = formValue.Elem()
	}
	t := reflect.TypeOf(formValue.Interface())
	for i := 0; i < t.NumField(); i++ {
		fieldName := t.Field(i).Name
		paramValue := formValue.FieldByName(fieldName)
		tag := t.Field(i).Tag.Get(pathKey)
		kind := paramValue.Kind()
		if kind == reflect.Struct {
			c.parsePathParams(paramValue.Addr().Interface())
		} else if tag != "" {
			if paramValue.IsValid() {
				switch kind {
				case reflect.Int:
					s := c.Param(tag)
					i64, _ := strconv.ParseInt(s, 10, 64)
					if i64 != 0 {
						paramValue.Set(reflect.ValueOf(i64))
					}

				case reflect.Uint:
					s := c.Param(tag)
					i64, _ := strconv.ParseInt(s, 10, 32)
					if i64 != 0 {
						paramValue.Set(reflect.ValueOf(uint(i64)))
					}

				case reflect.String:
					paramValue.Set(reflect.ValueOf(c.Param(tag)))
				}
			}
		}

		if paramValue.IsValid() {
			switch kind {
			case reflect.Ptr:
				pe := paramValue.Elem()
				if pe.Kind() == reflect.String &&
					pe.Interface() != nil &&
					pe.Interface().(string) == "" {
					paramValue.Set(reflect.ValueOf((*string)(nil)))
				}
			}
		}
	}
}

// HasUserContext has user context
func (c *Context) HasUserContext() bool {
	return c.Get(userContext) != nil
}

// GetContextID get context id
func (c *Context) GetContextID() string {
	return c.Request().Header.Get(echo.HeaderXRequestID)
}
