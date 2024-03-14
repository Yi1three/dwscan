package scan

import (
	"sync"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

type Task struct {
	Func   func(interface{}) interface{} // 函数变量
	Args   interface{}                   // 参数
	Result interface{}                   // 结果
}

func (t *Task) Process() {
	t.Result = t.Func(t.Args)
}

type WorkerPool struct {
	Tasks       []Task
	concurrency int // 线程池大小
	tasksChan   chan Task
	wg          sync.WaitGroup
	Results     []interface{} // 保存结果
	mu          sync.Mutex    // 保护 Results 的互斥锁
}

func (wp *WorkerPool) Worker() {
	for task := range wp.tasksChan {
		task.Process()

		// 保存结果
		wp.mu.Lock()
		wp.Results = append(wp.Results, task.Result)
		wp.mu.Unlock()

		wp.wg.Done()
	}
}

func (wp *WorkerPool) Run() {
	wp.tasksChan = make(chan Task, len(wp.Tasks))
	for i := 0; i < wp.concurrency; i++ {
		go wp.Worker()
	}

	wp.wg.Add(len(wp.Tasks))
	for _, task := range wp.Tasks {
		wp.tasksChan <- task
	}

	close(wp.tasksChan)
	wp.wg.Wait()
}

func InitWorkerPool(function func(interface{}) interface{}, args []interface{}, concurrency int) WorkerPool {

	pool := WorkerPool{
		Tasks:       make([]Task, len(args)),
		concurrency: concurrency,
	}

	for i, url := range args {

		pool.Tasks[i] = Task{
			Func: function,
			Args: url,
		}
	}
	return pool
}
func ExtractFormsFromHTML(html string) ([]map[string]interface{}, error) {
	var forms []map[string]interface{}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %v", err)
	}

	doc.Find("form").Each(func(i int, selection *goquery.Selection) {
		form := make(map[string]interface{})
		form["action"], _ = selection.Attr("action")

		selection.Find("input, textarea, select").Each(func(j int, input *goquery.Selection) {
			name, exists := input.Attr("name")
			if exists {
				value, exists := input.Attr("value")
				if exists {
					form[name] = value
				} else {
					form[name] = ""
				}
			}
		})

		forms = append(forms, form)
	})

	return forms, nil
}

func (wp *WorkerPool) FormExtractor(url string) Task {
	return Task{
		Func: func(arg interface{}) interface{} {
			resp, err := http.Get(arg.(string)) // 假设这是一个有效的GET请求
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}

		.forms, err := ExtractFormsFromHTML(string(body))
			if err != nil {
				return nil, err
			}

			return forms, nil
		},
		Args: url,
	}
}
