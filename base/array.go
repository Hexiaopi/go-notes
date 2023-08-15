package base

import "fmt"

func PrintArray() {
	var a [3]int
	var b = [...]int{1, 2, 3}
	var c = [...]int{1: 2, 2: 3}
	var d = [...]int{0, 2, 2: 3, 4}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
