package fanout

import (
	"fmt"
	"sync"
)


func worker(tasksCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		task, ok := <-tasksCh
		if !ok {
			return
		}
		fmt.Println("processing task", task)
	}
}

func pool(wg *sync.WaitGroup, workers, tasks int) {
	tasksCh := make(chan int)

	// consumer
	for i := 0; i < workers; i++ {
		go worker(tasksCh, wg)
	}
	//producer
	for i := 0; i < tasks; i++ {
		tasksCh <- i
	}

	close(tasksCh)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	go pool(&wg, 10, 5)
	wg.Wait()
}