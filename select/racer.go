package racer

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecondTimeout = time.Duration(10 * time.Second)

//Racer implements ConfigurableRacer for the real-world application.
func Racer(a, b string) (winner string, err error) {
	winner, err = ConfigurableRacer(a, b, tenSecondTimeout)
	return
}

//ConfigurableRacer is a function that takes two urls and returns the faster URL.
//Returns error if neither URL responds within 10seconds
//We can use this function to test, and Racer(implementation of this fn) for actual use
func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("Timed out waiting for %s and %s", a, b)
	}

}

//ping function takes a url and pings it thru a channel
func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch

}
