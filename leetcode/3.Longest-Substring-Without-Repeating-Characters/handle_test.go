package Longest_Substring_Without_Repeating_Characters

import "testing"

func Test_handleOne(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Example1", args: args{s: "abcabcbb"}, want: 3},
		{name: "Example2", args: args{s: "bbbb"}, want: 1},
		{name: "Example3", args: args{s: "pwwkew"}, want: 3},
		{name: "Example4", args: args{s: ""}, want: 0},
		{name: "Example5", args: args{s: " "}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handleOne(tt.args.s); got != tt.want {
				t.Errorf("handleOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handleTwo(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Example1", args: args{s: "abcabcbb"}, want: 3},
		{name: "Example2", args: args{s: "bbbb"}, want: 1},
		{name: "Example3", args: args{s: "pwwkew"}, want: 3},
		{name: "Example4", args: args{s: ""}, want: 0},
		{name: "Example5", args: args{s: " "}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handleTwo(tt.args.s); got != tt.want {
				t.Errorf("handleTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
