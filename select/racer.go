package racer

import "net/http"

//Racer is a function that takes two urls and returns the faster URL.
//Returns error if neither URL responds within 10seconds
func Racer(a, b string) string {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
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
