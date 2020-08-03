package concurrency

//WebsiteChecker is a function that takes a URL and checks if it up
type WebsiteChecker func(string) bool

//Result is a struct that captures the output of each goroutine into a channel
type result struct {
	string
	bool
}

//CheckWebsites function takes a list of URLs and returns a mapping of website to its status
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}
