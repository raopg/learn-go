package main

import (
	"fmt"
	"io"
	"net/http"
)

//Greet function uses the bytes buffer to greet someone
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

//Main function to use Greet function
// func main() {
// 	Greet(os.Stdout, "Prateek")
// }

//MyGreetHandler function is a HTTP handler for the Greet func.
func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Prateek")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreetHandler))
}
