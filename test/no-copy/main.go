package main

import "fmt"

type Demo struct {
}

func (*Demo) Lock()   {}
func (*Demo) Unlock() {}

func Copy(d Demo) {
}

func main() {
	d := Demo{}
	fmt.Printf("%+v", d)
	Copy(d)
	d2 := d
	fmt.Printf("%+v", d2)
}

// go vet
