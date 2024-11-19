package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	v := validator.New()
	v.RegisterTagNameFunc(SetTagName)
	v.RegisterValidation("default-name", SetDefaultName)
	v.RegisterValidation("idn-mobile-number", SetIDNMobileNumber)

	return &Validator{validate: v}
}

func (v *Validator) Validate(data interface{}) (err error) {
	err = v.validate.Struct(data)
	if err == nil {
		return
	}

	errorFields := err.(validator.ValidationErrors)
	errorField := errorFields[0]
	err = fmt.Errorf("invalid '%s' with value '%v'", errorField.Field(), errorField.Value())

	return
}
