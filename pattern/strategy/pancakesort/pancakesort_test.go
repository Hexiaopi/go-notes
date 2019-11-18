package pancakesort

import "testing"

func TestPancakeSort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	p := PancakeSort{}
	p.Sort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5 {
		t.Error("pancakesort() failed. Got", values, "Expected 1 2 3 4 5")
	}
}

func TestPancakeSort2(t *testing.T) {
	values := []int{5, 5, 3, 2, 1}
	p := PancakeSort{}
	p.Sort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 5 || values[4] != 5 {
		t.Error("pancakesort() failed. Got", values, "Expected 1 2 3 5 5")
	}
}

func TestPancakeSort3(t *testing.T) {
	values := []int{5}
	p := PancakeSort{}
	p.Sort(values)
	if values[0] != 5 {
		t.Error("pancakesort() failed. Got", values, "Expected 5")
	}
}
