package TwoSum

import (
	"reflect"
	"testing"
)

func Test_handleOne(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example1",
			args{nums: []int{2, 7, 11, 15}, target: 9},
			[]int{0, 1},
		},
		{"Example2",
			args{nums: []int{3,2,4}, target: 6},
			[]int{1,2},
		},
		{"Example3",
			args{nums: []int{3,3}, target: 6},
			[]int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handleOne(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handleOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handleTwo(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example1",
			args{nums: []int{2, 7, 11, 15}, target: 9},
			[]int{0, 1},
		},
		{"Example2",
			args{nums: []int{3,2,4}, target: 6},
			[]int{1,2},
		},
		{"Example3",
			args{nums: []int{3,3}, target: 6},
			[]int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handleTwo(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handleTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}