package base

import (
	"fmt"
)

func DeferDemo1() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

func DeferDemo2() {
	var i = 1
	defer fmt.Println(i)
	i = 2
}

func DeferDemo3() {
	var i = 1
	defer func() {
		fmt.Print(i)
	}()
	i = 2
}

func DeferDemo4() {
	var i = 1
	defer func(i int) {
		fmt.Print(i)
	}(i)
	i = 2
}

func DeferDemo5() (result int) {
	i := 1
	defer func() {
		result++
	}()
	return i
}

func DeferDemo6() int {
	i := 1
	defer func() {
		i++
	}()
	return i
}

func DeferDemo7() (result int) {
	defer func() {
		result++
	}()
	return 0
}

