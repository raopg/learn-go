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

	switch val.Kind() {
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.Struct:
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.String:
		for i := 0; i < val.Len(); i++ {
			fn(val.Field(i).String())
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
