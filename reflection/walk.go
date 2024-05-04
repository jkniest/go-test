package reflection

import (
	"reflect"
)

func walk(x any, fn func(input string)) {
	val := getValue(x)

	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walk(val.Field(i).Interface(), fn)
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walk(v.Interface(), fn)
			} else {
				break
			}
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walk(res.Interface(), fn)
		}
	case reflect.String:
		fn(val.String())
	}
}

func getValue(field any) reflect.Value {
	val := reflect.ValueOf(field)

	if val.Kind() == reflect.Pointer {
		return val.Elem()
	}

	return val
}
