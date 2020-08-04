package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("test fastURL is returned", func(t *testing.T) {
		slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(20 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
		}))

		fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		output := Racer(slowURL, fastURL)

		assertStringEquals(t, output, fastURL)

	})
}

func assertStringEquals(t *testing.T, output, expected string) {
	t.Helper()
	if output != expected {
		t.Errorf("Expected: %q\nGot %q", expected, output)
	}
}
