package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countDownCount = 3
const endMessage = "Go!"

//Sleeper encapsulates function of sleeper in two conditions - running & test
type Sleeper interface {
	Sleep()
}

//DefaultSleeper wraps the time.Sleep method in a Sleeper interface
type DefaultSleeper struct{}

//Sleep implementation of Default sleeper wraps time.Sleep
func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

//SpySleeper is a Mock of the Sleeper interface
type SpySleeper struct {
	Calls int
}

//Sleep implementation of SpySleeper updates the number of calls to it
func (s *SpySleeper) Sleep() {
	s.Calls++
}

//Countdown function counts down from 3 to 1 and then prints go
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countDownCount; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, endMessage)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
