package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("test fastURL is returned", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		output, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatal("Did not expect an error but found one")
		}

		assertStringEquals(t, output, fastURL)

	})

	t.Run("returns an error if Racer takes more than 10seconds", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("Expected an error but did not get one")
		}

	})
}

func assertStringEquals(t *testing.T, output, expected string) {
	t.Helper()
	if output != expected {
		t.Errorf("Expected: %q\nGot %q", expected, output)
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
