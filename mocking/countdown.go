package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countDownCount = 3
const endMessage = "Go!"

//Countdown function counts down from 3 to 1 and then prints go
func Countdown(out io.Writer) {
	for i := countDownCount; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprint(out, endMessage)
}

func main() {
	Countdown(os.Stdout)
}
