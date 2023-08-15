package base

import "testing"

func ExamplePrintString() {
	PrintString()
	// Output:
	// 11
}

func ExamplePrintStringCompare() {
	PrintStringCompare()
	// Output:
	// true
	// true
	// true
	// true
	// true
	// true
}

func ExamplePrintStringByte() {
	PrintStringByte("hello")
	// Output:
	// 0 h
	// 1 e
	// 2 l
	// 3 l
	// 4 o
}

func TestPrintStringByte(t *testing.T) {
	PrintStringByte("你好")
}

func ExamplePrintStringRune1() {
	PrintStringRune("hello")
	// Output:
	// 0 h
	// 1 e
	// 2 l
	// 3 l
	// 4 o
}

func ExamplePrintStringRune2() {
	PrintStringRune("你好")
	// Output:
	// 0 你
	// 3 好
}

func BenchmarkConcatStringByOperator(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringByOperator(sl)
	}
}

func BenchmarkConcatStringBySprintf(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringBySprintf(sl)
	}
}

func BenchmarkConcatStringByJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringByJoin(sl)
	}
}

func BenchmarkConcatStringByStringsBuilder(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringByStringsBuilder(sl)
	}
}

func BenchmarkConcatStringByStringsBuilderWithInitSize(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringByStringsBuilderWithInitSize(sl)
	}
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringByBytesBuffer(sl)
	}
}

func BenchmarkConcatStringByBytesBufferWithInitSize(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringByBytesBufferWithInitSize(sl)
	}
}
