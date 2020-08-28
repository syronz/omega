package helper

import (
	"errors"
	"reflect"
)

// Includes used for checking is item exist in the array or not
func Includes(arrayType interface{}, item interface{}) (result bool, err error) {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array && arr.Kind() != reflect.Slice {
		err = errors.New("invalid data-type")
		return
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true, nil
		}
	}

	return
}
