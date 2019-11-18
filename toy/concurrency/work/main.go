package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4
	taskLoad         = 10
)

var wg sync.WaitGroup

func main() {

	tasks := make(chan string, taskLoad)
	wg.Add(numberGoroutines)
	for i := 0; i < numberGoroutines; i++ {
		go worker(tasks, i+1)
	}
	for post := 0; post < taskLoad; post++ {
		tasks <- fmt.Sprintf("Task:%d", post+1)
	}
	close(tasks)
	wg.Wait()

}

func worker(tasks chan string, worker int) {
	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker: %d Shutting Down\n", worker)
			return
		}
		fmt.Printf("Worker: %d Started %s\n", worker, task)
		time.Sleep(time.Duration(rand.Int63n(100)) * time.Millisecond)
		fmt.Printf("Worker: %d Completed %s\n", worker, task)
	}
}
