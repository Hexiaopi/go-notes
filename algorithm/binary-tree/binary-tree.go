package binary_tree

import (
	"container/list"
	"fmt"
)

type Node struct {
	Left  *Node
	Data  interface{}
	Right *Node
}

func (n *Node) SetData(data interface{}) {
	n.Data = data
}

func (n *Node) SetLeft(node *Node) {
	n.Left = node
}

func (n *Node) SetRight(node *Node) {
	n.Right = node
}

func PreOrder(n *Node) {
	if n != nil {
		fmt.Printf("%v ", n.Data)
		PreOrder(n.Left)
		PreOrder(n.Right)
	}
}

func InOrder(n *Node) {
	if n != nil {
		InOrder(n.Left)
		fmt.Printf("%v ", n.Data)
		InOrder(n.Right)
	}
}

func AftOrder(n *Node) {
	if n != nil {
		AftOrder(n.Left)
		AftOrder(n.Right)
		fmt.Printf("%v ", n.Data)
	}
}

//层序遍历
func LayerOrder(n *Node) {
	queue := list.New()
	queue.PushBack(n)
	for queue.Len() > 0 {
		header := queue.Front()
		v := header.Value.(*Node)
		fmt.Print(v.Data," ")
		queue.Remove(header)
		if v.Left != nil {
			queue.PushBack(v.Left)
		}
		if v.Right != nil {
			queue.PushBack(v.Right)
		}
	}
}

func Print(n *Node) {
	if n != nil {
		fmt.Printf("%v", n.Data)
		if n.Left != nil || n.Right != nil {
			fmt.Printf("(")
			Print(n.Left)
			if n.Right != nil {
				fmt.Printf(",")
			}
			Print(n.Right)
			fmt.Printf(")")
		}
	}
}