package racer

import (
	"net/http"
	"time"
)

//Racer is a function that takes two urls and returns the faster URL.
//Returns error if neither URL responds within 10seconds
func Racer(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		winner = a
	} else {
		winner = b
	}

	return
}

//measureResponseTime func takes a URL and measures the time to get a response
func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
