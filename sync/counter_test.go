package counter

import "testing"

func TestCounter(t *testing.T) {
	t.Run("increment counter 3 times", func(t *testing.T) {
		c := Counter{}

		c.Inc()
		c.Inc()
		c.Inc()

		if c.Value() != 3 {
			t.Errorf("Expected 3, got %d", c.Value())
		}
	})
}
