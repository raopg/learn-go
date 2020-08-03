package concurrency

//WebsiteChecker is a function that takes a URL and checks if it up
type WebsiteChecker func(string) bool

//CheckWebsites function takes a list of URLs and returns a mapping of website to its status
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}
