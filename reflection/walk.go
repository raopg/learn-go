package reflection

import (
	"reflect"
)

//Person data type is used to test walk function
type Person struct {
	Name    string
	Profile Profile
}

//Profile data type is used to test walk function
type Profile struct {
	Age  int
	City string
}

func walk(data interface{}, fn func(input string)) {
	val := getValue(data)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		if field.Kind() == reflect.String {
			fn(field.String())
		}

		if field.Kind() == reflect.Struct {
			walk(field.Interface(), fn)
		}
	}
}

func getValue(data interface{}) reflect.Value {
	val := reflect.ValueOf(data)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
