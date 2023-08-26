package concurrency

import (
	"testing"
)

func TestFanIn(t *testing.T) {
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	in := echo(nums)
	c1 := square(in)
	c2 := square(in)

	for n := range merge(c1, c2) {
		t.Log(n)
	}
}

func TestFanOut(t *testing.T) {
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	in := echo(nums)
	c1 := odd(in)
	c2 := odd(in)
	c3 := odd(in)
	t.Log("c1")
	for n := range c1 {
		t.Log(n)
	}
	t.Log("c2")
	for n := range c2 {
		t.Log(n)
	}
	t.Log("c3")
	for n := range c3 {
		t.Log(n)
	}
}
