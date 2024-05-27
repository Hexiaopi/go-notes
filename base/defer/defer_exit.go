package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		fmt.Println("defer")
	}()
	fmt.Println("some thing...")
	os.Exit(1)
}
