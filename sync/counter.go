package counter

import "sync"

//Counter struct implements a simple concurrency-unsafe counter
type Counter struct {
	mu    sync.Mutex
	value int
}

//Inc will increment the current counter by 1
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

//Value will return the current value of the counter.
func (c *Counter) Value() int {
	return c.value
}
