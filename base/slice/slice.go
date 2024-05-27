package main

import "fmt"

func PrintSliceState() {
	var (
		a []int               // 指针为nil，len和cap为0
		b = []int{}           // 空切片，不等于nil，表示空的集合
		c = []int{1, 2, 3}    // 指向包含三个元素的数组：[1，2，3]。len和cap为3
		d = c[:2]             // 与c指向同一个数组，但是len为2，cap为3
		e = c[0:2:cap(c)]     // 与c指向同一个数组，但是len为2，cap为cap(c)的大小，即3
		f = c[:0]             // 与c指向同一个数组，但是len为0，cap为3
		g = make([]int, 3)    // 指向新的数组，包含三个元素：[0,0,0]。len和cap为3
		h = make([]int, 2, 3) // 指向新的数组，包含三个元素：[0,0,0]。len为2，cap为3
		i = make([]int, 0, 3) // 指向新的数组，包含三个元素：[0,0,0]。len为0，cap为3
	)
	fmt.Println(a, len(a), cap(a))
	fmt.Println(b, len(b), cap(b))
	fmt.Println(c, len(c), cap(c))
	fmt.Println(d, len(d), cap(d))
	fmt.Println(e, len(e), cap(e))
	fmt.Println(f, len(f), cap(f))
	fmt.Println(g, len(g), cap(g))
	fmt.Println(h, len(h), cap(h))
	fmt.Println(i, len(i), cap(i))
}

func PrintSliceAppend() {
	var s []int
	fmt.Println(len(s), cap(s)) // 0 0
	s = append(s, 1)
	fmt.Println(len(s), cap(s)) // 1 1
	s = append(s, 2)
	fmt.Println(len(s), cap(s)) // 2 2
	s = append(s, 3)
	fmt.Println(len(s), cap(s)) // 3 4
	s = append(s, 4)
	fmt.Println(len(s), cap(s)) // 4 4
	s = append(s, 5)
	fmt.Println(len(s), cap(s)) // 5 8
}

func SliceRise(s []int) {
	s = append(s, 0)
	for i := range s {
		s[i]++
	}
}

func SliceAppendConflict() {
	a := []int{1, 2, 3, 4, 5}
	b := a[1:3]
	b = append(b, 0)
	fmt.Println(b, len(b), cap(b)) // [2 3 0] 3 4
	fmt.Println(a, len(a), cap(a)) // [1 2 3 0 5] 5 5
}
