package golidation

import (
	"reflect"
	"regexp"
	"strconv"
)

type errors struct {
	Field string
	Error string
}

type Validator struct {
	Model    any
	Errors   []errors
	DevErrors []string
}

func (v *Validator) maxLength(max, fieldName string, field reflect.Value, OptionalError ...string) {

	i, err := strconv.Atoi(max)

	if err != nil {
		v.writeDevError(fieldName + ": maxL value should be number.")
		return
	}

	value, ok := field.Interface().(string)
	if !ok {
		v.writeDevError(fieldName + ": To check the maximum length, the field value should be string.")
		return
	}

	if len(value) > i {
		v.writeError(fieldName, fieldName+" field value is greater than the allowed maximum length.")
	}

}

func (v *Validator) minLength(min, fieldName string, field reflect.Value) {

	i, err := strconv.Atoi(min)

	if err != nil {
		v.writeDevError(fieldName + ": minL value should be number.")
		return
	}

	value, ok := field.Interface().(string)
	if !ok {
		v.writeDevError(fieldName + ": To check the minimum length, the field value should be string.")
		return
	}

	if len(value) < i {
		v.writeError(fieldName, fieldName+" field value is less than the allowed minimum length.")
	}
}

func (v *Validator) isEmail(fieldName string, field reflect.Value) {

	value, ok := field.Interface().(string)
	if !ok {
		v.writeDevError(fieldName + ": To check the Email, the field value should be string.")
		return
	}

	re := regexp.MustCompile(`[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}`)

	if !re.MatchString(value) {
		v.writeError(fieldName, fieldName+" field value is not a valid email.")
	}
}
