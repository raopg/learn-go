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

	numberOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		getField = val.Index
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field
	case reflect.String:
		fn(val.String())
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	}

	for i := 0; i < numberOfValues; i++ {
		walk(getField(i).Interface(), fn)
	}
}

func getValue(data interface{}) reflect.Value {
	val := reflect.ValueOf(data)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
