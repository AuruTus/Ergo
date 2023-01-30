package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func KeyGen(args ...any) string {
	var s string
	keys := make([]string, 0, 8)
	for _, a := range args {
		switch v := reflect.ValueOf(a); v.Kind() {
		case reflect.Pointer:
			if v.IsNil() {
				continue
			}
			s = fmt.Sprintf("0x%x", uint64(v.Pointer()))
		case reflect.String:
			s = v.String()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			ival := v.Int()
			if ival < 0 {
				ival = -ival
				s = fmt.Sprintf("m%d", ival)
				break
			}
			s = strconv.FormatInt(ival, 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			s = strconv.FormatUint(v.Uint(), 10)
		default:
			continue
		}
		keys = append(keys, s)
	}
	return strings.Join(keys, "-")
}
