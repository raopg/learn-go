package di

import (
	"fmt"
	"io"
	"os"
)

//Greet function uses the bytes buffer to greet someone
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

//Main function to use Greet function
func main() {
	Greet(os.Stdout, "Prateek")
}
