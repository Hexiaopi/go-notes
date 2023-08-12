package Longest_Palindromic_Substring

import "testing"

func Test_handleOne(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Example1", args: args{"babad"}, want: "bab"},
		{name: "Example2", args: args{"cbbd"}, want: "bb"},
		{name: "Example3", args: args{"a"}, want: "a"},
		{name: "Example4", args: args{"ac"}, want: "a"},
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
		want string
	}{
		{name: "Example1", args: args{"babad"}, want: "bab"},
		{name: "Example2", args: args{"cbbd"}, want: "bb"},
		{name: "Example3", args: args{"a"}, want: "a"},
		{name: "Example4", args: args{"ac"}, want: "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handleTwo(tt.args.s); got != tt.want {
				t.Errorf("handleTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handleThree(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Example1", args: args{"babad"}, want: "aba"},
		{name: "Example2", args: args{"cbbd"}, want: "bb"},
		{name: "Example3", args: args{"a"}, want: "a"},
		{name: "Example4", args: args{"ac"}, want: "c"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handleThree(tt.args.s); got != tt.want {
				t.Errorf("handleThree() = %v, want %v", got, tt.want)
			}
		})
	}
}