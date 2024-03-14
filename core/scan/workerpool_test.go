package scan

import (
	"fmt"
	"testing"
)

func TestTask_Process(t *testing.T) {
	urls := []interface{}{"url1", "url2", "url3"}

	pool := InitWorkerPool(ParseURL, urls, 5)

	pool.Run()

	// 打印任务结果
	for _, result := range pool.Results {
		fmt.Println("Task Result:", result)
	}
}

func ParseURL(url interface{}) interface{} {
	fmt.Println("Parsing URL:", url)
	return fmt.Sprintf("Result for URL: %s", url)
}
