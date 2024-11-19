package validator

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

// SetTagName tag name for validator
func SetTagName(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	if name == "-" {
		return ""
	}
	return name
}

// SetDefaultName validator for acceptable name format for mypertamina.
func SetDefaultName(fl validator.FieldLevel) bool {
	pattern := `^[A-Za-z. ']+$`
	stringNameRgx := regexp.MustCompile(pattern)
	return stringNameRgx.MatchString(fl.Field().String())
}

// SetIDNMobileNumber validator for idn mobile number.
func SetIDNMobileNumber(fl validator.FieldLevel) bool {
	pattern := `^(08)[0-9]+$`
	mobileNumberRgx := regexp.MustCompile(pattern)
	return mobileNumberRgx.MatchString(fl.Field().String())
}
