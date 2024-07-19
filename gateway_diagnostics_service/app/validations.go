package app

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var isHexadecimal validator.Func = func(fl validator.FieldLevel) bool {

	value, ok := fl.Field().Interface().(string)
	if ok {

		matched, _ := regexp.MatchString("^[a-fA-F0-9]{8}$", value)
		return matched
	}

	return false
}

var isSerialNumbers validator.Func = func(fl validator.FieldLevel) bool {

	values, ok := fl.Field().Interface().([]string)

	if ok {

		for _, value := range values {
			matched, _ := regexp.MatchString("^[a-fA-F0-9]{14}$", value)
			if matched == false {
				return false
			}
		}

		return true
	}

	return false
}
