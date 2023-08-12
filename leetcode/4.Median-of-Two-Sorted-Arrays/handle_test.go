package Median_of_Two_Sorted_Arrays

import "testing"

func Test_handleOne(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "Example1", args: args{nums1: []int{1, 3}, nums2: []int{2}}, want: 2},
		{name: "Example2", args: args{nums1: []int{1, 2}, nums2: []int{3,4}}, want: 2.5},
		{name: "Example3", args: args{nums1: []int{0, 0}, nums2: []int{0,0}}, want: 0},
		{name: "Example4", args: args{nums1: []int{}, nums2: []int{1}}, want: 1},
		{name: "Example5", args: args{nums1: []int{2}, nums2: []int{}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handleOne(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("handleOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
