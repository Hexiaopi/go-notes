package concurrency

import "testing"

func TestPipeline(t *testing.T) {
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for n := range square(odd(echo(nums))) {
		t.Log(n)
	}
}
