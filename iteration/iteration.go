package iteration

//Repeat takes the string to be repeated and a count. It returns a concatenated string
func Repeat(character string, times int) string {
	var repeated string // This is how you declare vars

	for i := 0; i < times; i++ { // This is how for loops are defined.
		repeated += character
	}

	return repeated
}
