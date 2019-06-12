package concurrency 

type result struct {
	string
	bool
}


// WebsiteChecker checks a website
type WebsiteChecker func(string) bool

// CheckWebsites does what it's name says
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string){
			resultChannel <- result{u, wc(u)}
		}(url)

		defer close(resultChannel)
		
	}
	for i :=0; i < len(urls); i++ {
		result := <- resultChannel
		results[result.string] = result.bool
	}

	// for result := range resultChannel {
	// 	results[result.string] = result.bool
	// }
	return results
}
