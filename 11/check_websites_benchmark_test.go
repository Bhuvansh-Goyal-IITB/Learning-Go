package concurrency

import (
	"testing"
	"time"
)

func slowWebsiteChecker(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	b.ResetTimer()

	for range b.N {
		CheckWebsites(slowWebsiteChecker, urls)
	}
}
