package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	round := make(chan int)

	wg.Add(2)

	go player("jack", round)
	go player("rose", round)

	round <- 1
	wg.Wait()
	fmt.Println("Terminating Program")
}

func player(name string, round chan int) {
	defer wg.Done()
	for {
		ball, ok := <-round
		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			close(round)
			return
		}
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		round <- ball
	}
}
