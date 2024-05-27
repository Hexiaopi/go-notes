package main

import (
	"fmt"
)

func ExampleDeferDemo1() {
	DeferDemo1()
	// Output:
	// 43210
}

func ExampleDeferDemo2() {
	DeferDemo2()
	// Output:
	// 1
}

func ExampleDeferDemo3() {
	DeferDemo3()
	// Output:
	// 2
}

func ExampleDeferDemo4() {
	DeferDemo4()
	// Output:
	// 1
}

func ExampleDeferDemo5() {
	fmt.Println(DeferDemo5())
	// Output:
	// 2
}

func ExampleDeferDemo6() {
	fmt.Println(DeferDemo6())
	// Output:
	// 1
}

func ExampleDeferDemo7() {
	fmt.Println(DeferDemo7())
	// Output:
	// 1
}
