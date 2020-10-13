package Add_Two_Numbers

import (
	"testing"
)

func TestPrintListNode(t *testing.T) {
	l1 := &ListNode{
		Val:  2,
		Next: nil,
	}
	l1.Next = &ListNode{
		Val:  4,
		Next: nil,
	}
	l1.Next.Next = &ListNode{
		Val:  3,
		Next: nil,
	}
	PrintListNode(l1)

	l2 := NewListNode([]int{4, 5, 6})
	PrintListNode(l2)
}

func TestSameListNode(t *testing.T) {
	l1 := NewListNode([]int{1, 2, 3})
	l2 := NewListNode([]int{1, 2, 3})
	l3 := NewListNode([]int{2, 3, 4})
	if !SameListNode(l1, l2) {
		t.Fatal("should same")
	}
	if SameListNode(l2, l3) {
		t.Fatal("should not same")
	}
}

func Test_handleOne(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"Example1",
			args{
				l1: NewListNode([]int{2, 4, 3}),
				l2: NewListNode([]int{5, 6, 4})},
			NewListNode([]int{7, 0, 8}),
		},
		{
			"Example2",
			args{
				l1: NewListNode([]int{0}),
				l2: NewListNode([]int{0})},
			NewListNode([]int{0}),
		},
		{
			"Example3",
			args{
				l1: NewListNode([]int{9, 9, 9, 9, 9, 9, 9}),
				l2: NewListNode([]int{9, 9, 9, 9})},
			NewListNode([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := handleOne(tt.args.l1, tt.args.l2)
			PrintListNode(got)
			if !SameListNode(got, tt.want) {
				t.Errorf("handleOne() = %v, want %v", *got, *tt.want)
			}
		})
	}
}



func Test_handleTwo(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"Example1",
			args{
				l1: NewListNode([]int{2, 4, 3}),
				l2: NewListNode([]int{5, 6, 4})},
			NewListNode([]int{7, 0, 8}),
		},
		{
			"Example2",
			args{
				l1: NewListNode([]int{0}),
				l2: NewListNode([]int{0})},
			NewListNode([]int{0}),
		},
		{
			"Example3",
			args{
				l1: NewListNode([]int{9, 9, 9, 9, 9, 9, 9}),
				l2: NewListNode([]int{9, 9, 9, 9})},
			NewListNode([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := handleTwo(tt.args.l1, tt.args.l2)
			PrintListNode(got)
			if !SameListNode(got, tt.want) {
				t.Errorf("handleOne() = %v, want %v", *got, *tt.want)
			}
		})
	}
}