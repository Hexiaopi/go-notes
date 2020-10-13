package Add_Two_Numbers

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(nums []int) *ListNode {
	var Node = &ListNode{
		Val:  0,
		Next: nil,
	}
	var result = Node
	for i := 0; i < len(nums); i++ {
		result.Val = nums[i]
		if i == len(nums)-1 {
			result.Next = nil
		} else {
			result.Next = &ListNode{0, nil}
			result = result.Next
		}
	}
	return Node
}

func PrintListNode(listNode *ListNode) {
	var node = listNode
	for node != nil {
		fmt.Println(node.Val)
		if node.Next != nil {
			node = node.Next
		} else {
			break
		}
	}
}

func SameListNode(l1 *ListNode, l2 *ListNode) bool {
	temp1, temp2 := l1, l2
	for temp1 != nil && temp2 != nil {
		if temp1.Val != temp2.Val {
			return false
		}
		temp1 = temp1.Next
		temp2 = temp2.Next
	}
	if temp1 != nil || temp2 != nil {
		return false
	}
	return true
}

func handleOne(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Node := l1
	l2Node := l2
	var next int = 0
	var sum int = 0
	var result = &ListNode{0, nil}
	var temp = result
	var left int = l1Node.Val
	var right int = l2Node.Val
	for l1Node != nil || l2Node != nil || next == 1 {
		sum = left + right + next
		if sum > 9 {
			sum = sum - 10
			next = 1
		} else {
			next = 0
		}
		temp.Val = sum
		//判断结果是否要新申请node，并移动指针
		if (l1Node != nil && l1Node.Next != nil) || (l2Node != nil && l2Node.Next != nil) || next == 1 {
			temp.Next = &ListNode{0, nil}
			temp = temp.Next
		}

		//计算下一次left，并移动指针
		if l1Node != nil {
			if l1Node.Next == nil {
				left = 0
				l1Node = nil
			} else {
				l1Node = l1Node.Next
				left = l1Node.Val
			}
		}
		//计算下一次right，并移动指针
		if l2Node != nil {
			if l2Node.Next == nil {
				right = 0
				l2Node = nil
			} else {
				l2Node = l2Node.Next
				right = l2Node.Val
			}
		}
	}
	return result
}

func handleTwo(l1 *ListNode,l2 *ListNode)*ListNode{
	head := &ListNode{Val: 0}
	n1, n2, carry, current := 0, 0, 0, head
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}
	return head.Next
}