package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("increment counter 3 times", func(t *testing.T) {
		c := Counter{}

		c.Inc()
		c.Inc()
		c.Inc()

		assertCounterVal(t, &c, 3)

	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		assertCounterVal(t, &counter, wantedCount)
	})
}

func assertCounterVal(t *testing.T, c *Counter, expected int) {
	t.Helper()
	if c.Value() != expected {
		t.Errorf("Expected %d, got %d", expected, c.Value())
	}
}
