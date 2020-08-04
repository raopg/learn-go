package racer

import "testing"

func TestRacer(t *testing.T) {
	t.Run("test fastURL is returned", func(t *testing.T) {
		fastURL := "https://google.com"
		slowURL := "https://raopg-chat-application.herokuapp.com"

		expected := fastURL

		output := Racer(fastURL, slowURL)

		assertStringEquals(t, output, expected)

	})
}

func assertStringEquals(t *testing.T, output, expected string) {
	t.Helper()
	if output != expected {
		t.Errorf("Expected: %q\nGot %q", expected, output)
	}
}
