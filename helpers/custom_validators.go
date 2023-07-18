package helpers

import (
	"errors"
	"reflect"
)

// Todo: Implement goplayground/validator
func IsArray(v interface{}, param string) error {
	slice := reflect.ValueOf(v)

	if slice.Kind() != reflect.Array {
		return errors.New("Powers must be an array of IDs")
	}

	if slice.Len() == 0 {
		return errors.New("Powers can't be empty")
	}

	return nil
}
