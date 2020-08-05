package reflection

import (
	"reflect"
)

func walk(data interface{}, fn func(input string)) {
	val := reflect.ValueOf(data)
	field := val.Field(0)
	fn(field.String())

}
