package Valid_Parentheses

import "testing"

func Test_handleOne(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Example1", args: args{"()"}, want: true},
		{name: "Example2", args: args{"()[]{}"}, want: true},
		{name: "Example3", args: args{"(]"}, want: false},
		{name: "Example4", args: args{"([)]"}, want: false},
		{name: "Example5", args: args{"{[]}"}, want: true},
		{name: "Example5", args: args{"{}[])("}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handleOne(tt.args.s); got != tt.want {
				t.Errorf("handleOne() = %v, want %v", got, tt.want)
			}
		})
	}
}