package tools

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const (
	pathErrorPattern        = "invalid path: %s"
	partialPathErrorPattern = "invalid path: %s in %s"
)

type PathError struct {
	seg  string
	path string
}

var _ error = (*PathError)(nil)

func (e *PathError) Error() string {
	if len(e.seg) != 0 {
		return fmt.Sprintf(partialPathErrorPattern, e.seg, e.path)
	}
	return fmt.Sprintf(pathErrorPattern, e.path)
}

// getJsonField is a tool to fetch the underlying value of the given path
func getJsonField[M ~map[string]any](m M, path string) (reflect.Value, error) {
	p := strings.FieldsFunc(path, func(r rune) bool {
		switch r {
		case '.', '[', ']':
			return true
		}
		return false
	})
	nd := reflect.ValueOf(m)

	for range p {
		switch nd.Kind() {
		case reflect.Map:
			nd = nd.MapIndex(reflect.ValueOf(p[0]))
		case reflect.Struct:
			nd = nd.FieldByName(p[0])
		case reflect.Pointer, reflect.Interface:
			nd = nd.Elem()
			continue
		case reflect.Array, reflect.Slice, reflect.String:
			i, err := strconv.Atoi(p[0])
			if err != nil || i >= nd.Len() {
				return reflect.Value{}, &PathError{seg: p[0], path: path}
			}
			nd = nd.Index(i)
		default:
			return reflect.Value{}, &PathError{seg: p[0], path: path}
		}
		if !nd.IsValid() {
			return reflect.Value{}, &PathError{seg: p[0], path: path}
		}
		p = p[1:]
		for nd.Kind() == reflect.Interface {
			nd = nd.Elem()
		}
	}

	if len(p) != 0 {
		return reflect.Value{}, &PathError{path: path}
	}
	// return nd.Elem(), nil
	return nd, nil
}

func GetJsonField(j []byte, path string) (val reflect.Value, err error) {
	if r := WithRecover(func() {
		m := make(map[string]any)
		if err = json.Unmarshal(j, &m); err != nil {
			return
		}
		val, err = getJsonField(m, path)
	}); r != nil {
		err = r
	}
	return
}
