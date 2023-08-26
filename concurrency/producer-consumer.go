package concurrency

import "fmt"

// Producer: 生产factor倍数
func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

// Consumer: 消费打印队列
func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
