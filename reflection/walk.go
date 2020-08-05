package reflection

import (
	"reflect"
)

func walk(data interface{}, fn func(input string)) {
	val := reflect.ValueOf(data)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String {
			fn(field.String())
		}
	}

}
