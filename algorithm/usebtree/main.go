package main

import (
	"DataStructuresAndAlgorithm/btree"
	"fmt"
)

func main() {
	//创建根节点
	root := btree.NewNode(nil, nil)
	var it btree.Initer
	it = root
	it.SetData("root node")
	//创建左子树
	a := btree.NewNode(nil, nil)
	a.SetData("left node")
	al := btree.NewNode(nil, nil)
	al.SetData(100)
	ar := btree.NewNode(nil, nil)
	ar.SetData(3.14)
	a.Left = al
	a.Right = ar
	//创建右子树
	b := btree.NewNode(nil, nil)
	b.SetData("right node")
	root.Left = a
	root.Right = b
	var ioperater btree.Operater
	ioperater = root
	ioperater.PrintBT()
	fmt.Println()
	var iorder btree.Order
	iorder = root
	iorder.PreOrder()
	fmt.Println()
	iorder.InOrder()
	fmt.Println()
	iorder.PostOrder()
}
