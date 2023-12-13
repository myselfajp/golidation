package golidation

import (
	"fmt"
	"reflect"
	"strings"
)

func (v *Validator) writeError(fieldName, errorText string) {
	err := errors{Field: fieldName, Error: errorText}
	v.Errors = append(v.Errors, err)
}

func (v *Validator) writeDevError(errorText string) {
	v.DevErrors = append(v.DevErrors, errorText)
}

func (validator *Validator) Check() {

	t := reflect.TypeOf(validator.Model)

	if t.Kind() != reflect.Struct {
		validator.writeDevError(fmt.Sprintf("showTags: expected a struct, got %v", t.Kind()))
		return
	}

	for i := 0; i < t.NumField(); i++ {

		var (
			field      = t.Field(i)
			fieldValue = reflect.ValueOf(validator.Model).Field(i)
			isEmpty    = reflect.DeepEqual(fieldValue.Interface(), reflect.Zero(fieldValue.Type()).Interface())
		)

		tag := field.Tag.Get("validate")
		required := strings.Contains(tag, "required")

		if !required && isEmpty {
			continue
		} else if required && isEmpty {
			validator.writeError(field.Name, field.Name+" is required field and can not be empty.")
			continue
		}

		if tag != "" {
			tagParts := strings.Split(tag, ",")
			for _, part := range tagParts {
				keyValue := strings.Split(part, "=")

				if len(keyValue) == 2 {
					key, value := keyValue[0], keyValue[1]
					switch key {
					case "maxL":
						validator.maxLength(value, field.Name, fieldValue)
					case "minL":
						validator.minLength(value, field.Name, fieldValue)
					case "email":
						validator.isEmail(field.Name, fieldValue)
					default:
						validator.writeDevError(key + "is not a valid tag")
					}
				} else if len(keyValue) == 1 {
					key := keyValue[0]
					switch key {
					case "email":
						validator.isEmail(field.Name, fieldValue)
					case "required":
					default:
						validator.writeDevError(key + "is not a valid tag")
					}
				}
			}
		}
	}
}
