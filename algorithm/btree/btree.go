package btree

import (
	"fmt"
)

type Node struct {
	Left  *Node
	Data  interface{}
	Right *Node
}

//用于初始化的接口
type Initer interface {
	SetData(data interface{})
}

//输出、深度计算、叶子统计基本操作接口
type Operater interface {
	PrintBT()
	Depth() int
	LeafCount() int
}

//前序遍历、中序遍历、后续遍历
type Order interface {
	PreOrder()
	InOrder()
	PostOrder()
}

func (n *Node) SetData(data interface{}) {
	n.Data = data
}

func (n *Node) PrintBT() {
	PrintBT(n) //调用底层函数
}

func (n *Node) Depth() int {
	return Depth(n) //调用底层函数
}

func (n *Node) LeafCount() int {
	return LeafCount(n) //调用底层函数
}

func (n *Node) PreOrder() {
	PreOrder(n) //调用底层函数
}

func (n *Node) InOrder() {
	InOrder(n) //调用底层函数
}

func (n *Node) PostOrder() {
	PostOrder(n) //调用底层函数
}

//调用底层函数
func NewNode(left, right *Node) *Node {
	return &Node{left, nil, right}
}

func PrintBT(n *Node) {
	if n != nil {
		fmt.Printf("%v", n.Data)
		if n.Left != nil || n.Right != nil {
			fmt.Printf("(")
			PrintBT(n.Left)
			if n.Right != nil {
				fmt.Printf(",")
			}
			PrintBT(n.Right)
			fmt.Printf(")")
		}
	}
}

func Depth(n *Node) int {
	var depleft, depright int
	if n == nil {
		return 0
	} else {
		depleft = Depth(n.Left)
		depright = Depth(n.Right)
		if depleft > depright {
			return depleft + 1
		} else {
			return depright + 1
		}
	}
}

func LeafCount(n *Node) int {
	if n == nil {
		return 0
	} else if (n.Left == nil) && (n.Right == nil) {
		return 1
	} else {
		return (LeafCount(n.Left) + LeafCount(n.Right))
	}
}

func PreOrder(n *Node) {
	if n != nil {
		fmt.Printf(" %v ", n.Data)
		PreOrder(n.Left)
		PreOrder(n.Right)
	}
}

func InOrder(n *Node) {
	if n != nil {
		PreOrder(n.Left)
		fmt.Printf(" %v ", n.Data)
		PreOrder(n.Right)
	}
}

func PostOrder(n *Node) {
	PreOrder(n.Left)
	PreOrder(n.Right)
	fmt.Printf(" %v ", n.Data)
}
