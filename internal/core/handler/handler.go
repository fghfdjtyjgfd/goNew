package handler

import (
	"net/http"
	"reflect"

	"basic-api/internal/core/context"

	"github.com/labstack/echo/v4"
)

// ResponseObject handle response object
func ResponseObject(c echo.Context, fn interface{}, form interface{}) error {
	cc := c.(*context.Context)
	err := cc.BindAndValidate(form)
	if err != nil {
		cc.Log.Error(err)
		return err
	}
	out := reflect.ValueOf(fn).Call([]reflect.Value{
		reflect.ValueOf(c),
		reflect.ValueOf(form),
	})
	errObj := out[1].Interface()
	if errObj != nil {
		return errObj.(error)
	}
	return cc.JSONResponse(http.StatusOK, out[0].Interface())
}

// ResponseObjectWithoutRequest handle response object without request
func ResponseObjectWithoutRequest(c echo.Context, fn interface{}) error {
	cc := c.(*context.Context)
	out := reflect.ValueOf(fn).Call([]reflect.Value{
		reflect.ValueOf(c),
	})

	errObj := out[1].Interface()
	if errObj != nil {
		return errObj.(error)
	}

	return cc.JSONResponse(http.StatusOK, out[0].Interface())
}

// ResponseSuccess handle response success
func ResponseSuccess(c echo.Context, fn interface{}, form interface{}) error {
	cc := c.(*context.Context)
	err := cc.BindAndValidate(form)
	if err != nil {
		cc.Log.Error(err)
		return err
	}

	out := reflect.ValueOf(fn).Call([]reflect.Value{
		reflect.ValueOf(c),
		reflect.ValueOf(form),
	})

	errObj := out[0].Interface()
	if errObj != nil {
		return errObj.(error)
	}
	return cc.JSONResponse(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

// ResponseSuccessWithoutRequest handle response success without request
func ResponseSuccessWithoutRequest(c echo.Context, fn interface{}) error {
	cc := c.(*context.Context)
	out := reflect.ValueOf(fn).Call([]reflect.Value{
		reflect.ValueOf(c),
	})

	errObj := out[0].Interface()
	if errObj != nil {
		return errObj.(error)
	}

	return cc.JSONResponse(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}
