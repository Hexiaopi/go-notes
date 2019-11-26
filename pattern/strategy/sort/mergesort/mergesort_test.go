package mergesort

import "testing"

func TestMergeSort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	m := MergeSort{Low:0,High:len(values)-1}
	m.Sort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5 {
		t.Error("mergesort() failed.Got", values, "Expected 1 2 3 4 5")
	}
}

func TestMergeSort2(t *testing.T) {
	values := []int{5, 5, 3, 2, 1}
	m := MergeSort{Low:0,High:len(values)-1}
	m.Sort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 5 || values[4] != 5 {
		t.Error("mergesort() failed.Got", values, "Expected 1 2 3 5 5")
	}
}

func TestMergeSort3(t *testing.T) {
	values := []int{5}
	m := MergeSort{Low:0,High:len(values)-1}
	m.Sort(values)
	if values[0] != 5 {
		t.Error("mergesort() failed.Got", values, "Expected 5")
	}
}
