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

//ConfigurableSleeper is a variant of Sleeper interface that allows you to configure the time
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

//Sleep method of ConfigurableSleeper
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

//SpyTime is a struct that holds the sleep timer
type SpyTime struct {
	durationSlept time.Duration
}

//Sleep implementation of SpyTime
func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

//Constants to recongize ops
const write = "Write"
const sleep = "Sleep"

//CountdownOperationsSpy checks the order of operations within the countdown function
type CountdownOperationsSpy struct {
	Calls []string
}

//Sleep implemented by the CountdownOperationsSpy struct
func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

//Write function implemented by CountdownOperations struct
func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
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
