package reflection

func walk(data interface{}, fn func(input string)) {
	fn("Hello there")
}
