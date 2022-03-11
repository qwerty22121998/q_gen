package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var ErrNilValue = errors.New("nil value")

func Extract(i interface{}, path string) (interface{}, error) {
	if path == "" {
		return i, nil
	}
	return extract(i, "", strings.Split(path, "."))
}

func ExtractPath(i interface{}, path[]string) (interface{}, error) {
	return extract(i, "", path)
}

func extract(s interface{}, key string, path []string) (interface{}, error) {
	if s == nil {
		return nil, ErrNilValue
	}
	if len(path) == 0 {
		return s, nil
	}
	ref := reflect.ValueOf(s)
	now := ref.FieldByName(path[0])
	if !now.IsValid() {
		return nil, fmt.Errorf("field %v.%v does not exist", key, path[0])
	}
	if len(path) > 1 {
		return extract(now.Interface(), key+path[0], path[1:])
	}
	return now.Interface(), nil
}
