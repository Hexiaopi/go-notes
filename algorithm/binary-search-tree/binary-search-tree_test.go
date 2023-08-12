package binary_search_tree

import (
	"fmt"
	"testing"
)

func ExampleNode_Add() {
	var node = &Node{
		Left:  nil,
		Data:  10,
		Right: nil,
	}
	node.Add(7)
	node.Add(8)
	node.Add(13)
	node.Add(14)
	node.Print()
	// Output:
	// 10(7(,8),13(,14))
}

func TestNode_Find(t *testing.T) {
	var node = &Node{
		Left:  nil,
		Data:  10,
		Right: nil,
	}
	node.Add(7)
	node.Add(8)
	node.Add(13)
	node.Add(14)
	result := node.Find(13)
	if result == nil || result.Data != 13 {
		t.Fatal("find fail")
	}
	t.Log(result.Data, result.Left, result.Right)

	result = node.Find(11)
	if result != nil {
		t.Fatal("should be nil")
	}
}

func TestNode_FindParent(t *testing.T) {
	var node = &Node{
		Left:  nil,
		Data:  10,
		Right: nil,
	}
	node.Add(7)
	node.Add(8)
	node.Add(13)
	node.Add(14)
	result := node.FindParent(14)
	if result == nil || result.Data != 13 {
		t.Fatal("find parent fail")
	}
	t.Log(result.Data, result.Left, result.Right)
	result = node.FindParent(10)
	if result != nil {
		t.Fatal("should be nil")
	}
	t.Log(result)
	var nodeNil *Node
	result = nodeNil.FindParent(1)
	if result != nil {
		t.Fatal("should be nil")
	}
	t.Log(result)
}

func ExampleNode_Delete() {
	var node = &Node{
		Left:  nil,
		Data:  10,
		Right: nil,
	}
	node.Add(7)
	node.Add(8)
	node.Add(13)
	node.Add(14)
	node.Print()
	fmt.Println()
	node = node.Delete(7)
	node.Print()
	fmt.Println()
	node = node.Delete(14)
	node.Print()
	fmt.Println()
	node = node.Delete(8)
	node.Print()
	fmt.Println()
	node = node.Delete(10)
	node.Print()
	// Output:
	// 10(7(,8),13(,14))
	// 10(8,13(,14))
	// 10(8,13)
	// 10(,13)
	// 13
}

func ExampleDeleteNodeRecursion() {
	var node = &Node{
		Left:  nil,
		Data:  10,
		Right: nil,
	}
	node.Add(7)
	node.Add(6)
	node.Add(9)
	node.Add(8)
	node.Add(13)
	node.Add(14)
	node.Print()
	fmt.Println()
	node = node.DeleteNodeRecursion(7)
	node.Print()
	// Output:
	// 10(7(6,9(8)),13(,14))
	// 10(8(6,9),13(,14))

}

func ExampleDeleteNodeIteration() {
	var node = &Node{
		Left:  nil,
		Data:  10,
		Right: nil,
	}
	node.Add(7)
	node.Add(6)
	node.Add(9)
	node.Add(8)
	node.Add(13)
	node.Add(14)
	node.Print()
	fmt.Println()
	node = node.DeleteNodeIteration(7)
	node.Print()
	// Output:
	// 10(7(6,9(8)),13(,14))
	// 10(8(6,9),13(,14))
}
