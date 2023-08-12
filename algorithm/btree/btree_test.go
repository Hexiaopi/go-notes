package btree

import (
	"fmt"
	"testing"
)

func TestPrintBT(t *testing.T) {
	//创建根节点
	root := NewNode(nil, nil)
	var it Initer
	it = root
	it.SetData("root node")
	//创建左子树
	a := NewNode(nil, nil)
	a.SetData("left node")
	al := NewNode(nil, nil)
	al.SetData(100)
	ar := NewNode(nil, nil)
	ar.SetData(3.14)
	a.Left = al
	a.Right = ar
	//创建右子树
	b := NewNode(nil, nil)
	b.SetData("right node")
	root.Left = a
	root.Right = b
	var ioperater Operater
	ioperater = root
	ioperater.PrintBT()
	fmt.Println()
	var iorder Order
	iorder = root
	iorder.PreOrder()
	fmt.Println()
	iorder.InOrder()
	fmt.Println()
	iorder.PostOrder()
}