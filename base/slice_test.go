package base

import "testing"

func ExamplePrintSliceState() {
	PrintSliceState()
	// Output:
	// [] 0 0
	// [] 0 0
	// [1 2 3] 3 3
	// [1 2] 2 3
	// [1 2] 2 3
	// [] 0 3
	// [0 0 0] 3 3
	// [0 0] 2 3
	// [] 0 3
}

func ExamplePrintSliceAppend() {
	PrintSliceAppend()
	// Output:
	// 0 0
	// 1 1
	// 2 2
	// 3 4
	// 4 4
	// 5 8
}

func TestSliceRise(t *testing.T) {
	s1 := []int{1, 2}
	s2 := s1
	s2 = append(s2, 3)
	SliceRise(s1)
	SliceRise(s2)
	t.Log(s1, len(s1), cap(s1))
	t.Log(s2, len(s2), cap(s2))
}

func ExampleSliceAppendConflict() {
	SliceAppendConflict()
	// Output:
	// [2 3 0] 3 4
	// [1 2 3 0 5] 5 5
}

const size = 10000

func BenchmarkSliceInitWithoutCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s1 := make([]int, 0)
		for i := 0; i < size; i++ {
			s1 = append(s1, i)
		}
	}
}

func BenchmarkSliceInitWithCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s1 := make([]int, 0, size)
		for i := 0; i < size; i++ {
			s1 = append(s1, i)
		}
	}
}
