package main

import (
	"net/http"
	"time"
)

//Racer is a function that takes two urls and returns the faster URL.
//Returns error if neither URL responds within 10seconds
func Racer(a, b string) (winner string) {
	//Start 'a' timer, get 'a' and capture time since.
	startA := time.Now()
	http.Get(a)
	aDuration := time.Since(startA)

	//Start 'b' timer, get 'b' and capture time since.
	startB := time.Now()
	http.Get(b)
	bDuration := time.Since(startB)

	if aDuration < bDuration {
		winner = a
	} else {
		winner = b
	}

	return
}
