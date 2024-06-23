package main

import (
	"fmt"
	"unsafe"
)

type S1 struct {
	a int8
	b int64
	c int8
	d int32
	e int16
}

type S2 struct {
	a int8
	c int8
	e int16
	d int32
	b int64
}

func main() {
	fmt.Println("Size of S1:", unsafe.Sizeof(S1{}), "bytes")
	fmt.Println(unsafe.Offsetof(S1{}.a), unsafe.Offsetof(S1{}.b), unsafe.Offsetof(S1{}.c), unsafe.Offsetof(S1{}.d), unsafe.Offsetof(S1{}.e))
	fmt.Println("Size of S2:", unsafe.Sizeof(S2{}), "bytes")
	fmt.Println(unsafe.Offsetof(S2{}.a), unsafe.Offsetof(S2{}.c), unsafe.Offsetof(S2{}.e), unsafe.Offsetof(S2{}.d), unsafe.Offsetof(S2{}.b))
}
