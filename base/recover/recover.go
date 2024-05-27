package main

import "fmt"

func PrintRecover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("a")
	panic("some thing")
	fmt.Println("b")
}

func FailRecoverWithoutDefer() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("a")
	panic("some thing")
	fmt.Println("b")
}

func FailRecoverWithoutLevelDefer() {
	defer func() {
		func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
	}()
	fmt.Println("a")
	panic("some thing")
	fmt.Println("b")
}
