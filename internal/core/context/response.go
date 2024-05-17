package context

// JSONResponse custom response
func (c *Context) JSONResponse(code int, i interface{}) error {

	return c.JSON(code, i)
}
