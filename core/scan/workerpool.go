package scan

import "fmt"
import "time"

// 这里是 worker，我们将并发执行多个 worker。
// worker 将从 `jobs` 通道接收任务，并且通过 `results` 发送对应的结果。
// 我们将让每个任务间隔 1s 来模仿一个耗时的任务。
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}
