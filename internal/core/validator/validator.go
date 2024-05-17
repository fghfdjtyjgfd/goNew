package validator

import (
	"fmt"
	"log"
	"regexp"

	"gopkg.in/go-playground/validator.v9"
)

// Validator validator
type Validator struct {
	validator *validator.Validate
}

// New new
func New() *Validator {
	v := &Validator{
		validator: validator.New(),
	}

	v.customValidator()
	return v
}

func (cv *Validator) customValidator() {
	_ = cv.validator.RegisterValidation("sort", func(fl validator.FieldLevel) bool {
		return validSort(fl.Field().String())
	})
}

// Validate validator
func (cv *Validator) Validate(i interface{}) error {
	log.Println("interface: ", i)
	return cv.validator.Struct(i)
}

// Var var
func (cv *Validator) Var(field interface{}, tag string) error {
	log.Println(fmt.Sprintf("field: %v, tag: %v", field, tag))
	return cv.validator.Var(field, tag)
}

func validSort(s string) bool {
	if s == "" {
		return true
	}
	var re = regexp.MustCompile(`(?m)^[a-z0-9_.]+$`)
	return re.MatchString(s)
}
