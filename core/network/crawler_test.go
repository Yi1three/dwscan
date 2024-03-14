package network

import (
	"dwscan/core/scan"
	"fmt"
	"testing"
)

func TestCrawl(t *testing.T) {
	request := NewGETRequest("http://www.zhattatey.top", false, nil)
	request2 := NewGETRequest("http://www.bilibili.com", false, nil)

	args := []interface{}{request, request2}

	pool := scan.InitWorkerPool(Crawl, args, 5)
	pool.Run()

	for _, result := range pool.Results {
		fmt.Println("Task Result:", result)
	}
}
