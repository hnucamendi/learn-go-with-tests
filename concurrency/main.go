package main

type WebsiteChecker func(string) bool
type result struct {
	url     string
	isValid bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultsChannel <- result{url: u, isValid: wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultsChannel
		results[result.url] = result.isValid
	}

	return results
}
