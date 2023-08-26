package concurrency

import (
	"testing"
	"time"
)

func TestProducerConsumer(t *testing.T) {
	ch := make(chan int, 10)
	go Producer(3, ch)
	go Producer(5, ch)
	go Consumer(ch)
	time.Sleep(1 * time.Second)
}
